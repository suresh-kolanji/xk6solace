package isolace

import (
	"encoding/json"
	"errors"

	"github.com/dop251/goja"
	"github.com/sirupsen/logrus"
	"go.k6.io/k6/js/common"
	"go.k6.io/k6/js/modules"
	"solace.dev/go/messaging"
	"solace.dev/go/messaging/pkg/solace"
	"solace.dev/go/messaging/pkg/solace/config"
)

type Solaceconfig struct {
	Host     string `json:"host"`
	Vpn      string `json:"vpn"`
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type Solace struct {
	vu     modules.VU
	logger *logrus.Logger
}

func NewSolace(vu modules.VU, logger *logrus.Logger) *Solace {
	return &Solace{vu: vu, logger: logger}
}

func (s *Solace) SolaceConnection(call goja.ConstructorCall) *goja.Object {
	runtime := s.vu.Runtime()
	var solaceConfig *Solaceconfig
	if len(call.Arguments) == 0 {
		common.Throw(runtime, errors.New("Custom Error"))
	}

	if params, ok := call.Argument(0).Export().(map[string]interface{}); ok {
		if b, err := json.Marshal(params); err != nil {
			common.Throw(runtime, err)
		} else {
			if err = json.Unmarshal(b, &solaceConfig); err != nil {
				common.Throw(runtime, err)
			}
		}
	}
	// Configuration parameters
	brokerConfig := config.ServicePropertyMap{
		config.TransportLayerPropertyHost:                solaceConfig.Host,
		config.ServicePropertyVPNName:                    solaceConfig.Vpn,
		config.AuthenticationPropertySchemeBasicPassword: solaceConfig.Password,
		config.AuthenticationPropertySchemeBasicUserName: solaceConfig.UserName,
	}
	s.logger.Println(solaceConfig)
	// Skip certificate validation
	messagingService, err := messaging.NewMessagingServiceBuilder().
		FromConfigurationProvider(brokerConfig).
		WithTransportSecurityStrategy(config.NewTransportSecurityStrategy().WithoutCertificateValidation()).
		Build()
	if err != nil {
		panic(err)
	}

	// Connect to the messaging serice
	if err := messagingService.Connect(); err != nil {
		panic(err)
	}

	s.logger.Print("Connected to the broker? ", messagingService.IsConnected())
	return runtime.ToValue(messagingService).ToObject(runtime)
}

func (s *Solace) CloseSolaceConnection(ms solace.MessagingService) {
	ms.Disconnect()
}
