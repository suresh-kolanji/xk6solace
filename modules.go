package ikeaxk6

import (
	"github.com/dop251/goja"
	"github.com/sirupsen/logrus"
	dsolace "github.com/suresh-kolanji/xk6solace/pkg/isolace/directpublisher"
	isolace "github.com/suresh-kolanji/xk6solace/pkg/isolace/messagingservice"
	"go.k6.io/k6/js/common"
	"go.k6.io/k6/js/modules"
)

// logger is used globally fro ikeaxk6 module
var logger *logrus.Logger

type Ikeaxk6 struct {
	vu            modules.VU
	metrics       ikeaxk6Metrics
	exports       *goja.Object
	imsolace      *isolace.Solace
	directpublish *dsolace.DirectPublish
}
type RootModule struct{}
type ModuleInstance struct{ *Ikeaxk6 }

func init() {
	logger = logrus.New()
	modules.Register("k6/x/ikeaxk6", new(RootModule))
}

var (
	_ modules.Instance = &ModuleInstance{}
	_ modules.Module   = &RootModule{}
)

// NewModuleInstance creates a new instance of the ikeaxk6 module.
func (*RootModule) NewModuleInstance(virtualUser modules.VU) modules.Instance {
	runtime := virtualUser.Runtime()
	metrics, err := registerMetrics(virtualUser)
	if err != nil {
		common.Throw(virtualUser.Runtime(), err)
	}
	msolace := isolace.NewSolace(virtualUser, logger)
	directPublisher := dsolace.NewDirectPublish(virtualUser, logger)
	// Create a new ikeaxk6 module.
	moduleInstance := &ModuleInstance{
		Ikeaxk6: &Ikeaxk6{
			vu:            virtualUser,
			metrics:       metrics,
			exports:       runtime.NewObject(),
			imsolace:      msolace,
			directpublish: directPublisher,
		},
	}
	mustExport := func(name string, value interface{}) {
		if err := moduleInstance.exports.Set(name, value); err != nil {
			common.Throw(runtime, err)
		}
	}

	//Exporting the solace object to JS From where we can invoke other solace modules
	mustExport("Messagingservice", moduleInstance.Ikeaxk6.imsolace.SolaceConnection)
	mustExport("DirectPublish", moduleInstance.Ikeaxk6.directpublish.SolaceDirectPublisher)
	mustExport("DisconnectMessService", moduleInstance.Ikeaxk6.imsolace.CloseSolaceConnection)
	mustExport("TerminateDirectPublish", moduleInstance.Ikeaxk6.directpublish.DisconnectPublisher)

	// This causes the struct fields to be exported to the native (camelCases) JS code.

	virtualUser.Runtime().SetFieldNameMapper(goja.TagFieldNameMapper("json", true))
	return moduleInstance
}

// Exports returns the exports of the ikeaxk6 module, which are the functions
// that can be called from the JS code.
func (m *ModuleInstance) Exports() modules.Exports {
	return modules.Exports{
		Default: m.Ikeaxk6.exports,
	}
}
