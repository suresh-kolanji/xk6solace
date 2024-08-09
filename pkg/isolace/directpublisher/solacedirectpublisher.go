package isolace

import (
	"errors"

	"github.com/dop251/goja"
	"github.com/sirupsen/logrus"
	"go.k6.io/k6/js/common"
	"go.k6.io/k6/js/modules"
	"solace.dev/go/messaging/pkg/solace"
	"solace.dev/go/messaging/pkg/solace/resource"
)

type DirectPublish struct {
	vu     modules.VU
	logger *logrus.Logger
}

func NewDirectPublish(vu modules.VU, logger *logrus.Logger) *DirectPublish {
	return &DirectPublish{vu: vu, logger: logger}
}

func (p *DirectPublish) SolaceDirectPublisher(call goja.ConstructorCall) *goja.Object {
	runtime := p.vu.Runtime()
	if len(call.Arguments) == 0 {
		common.Throw(runtime, errors.New("Custom Error no enough argument"))
		p.logger.Println("Not Enough Arg")
	}
	var messagingSevice interface{}
	var tmptopic interface{}
	var tmpPayload interface{}
	if params, ok := call.Argument(0).Export().(map[string]interface{}); ok {
		messagingSevice = params["connection"]
		tmptopic = params["topic"]
		tmpPayload = params["message"]
	}
	ms := messagingSevice.(solace.MessagingService)
	topicId := tmptopic.(string)
	stringPayload := tmpPayload.(string)
	directPublisher, builderErr := ms.CreateDirectMessagePublisherBuilder().Build()
	if builderErr != nil {
		panic(builderErr)
		p.logger.Error(builderErr)
	}
	startErr := directPublisher.Start()
	if startErr != nil {
		panic(startErr)
		p.logger.Error(startErr)
	}

	p.logger.Println("Direct Publisher running? ", directPublisher.IsRunning())

	p.logger.Println("\n===Interrupt (CTR+C) to stop publishing===\n")

	//	  Prepare outbound message payload and body
	messageBuilder := ms.MessageBuilder().
		WithProperty("application", "samples").
		WithProperty("language", "go")

		//	 Run forever until an interrupt signal is received
	message, err := messageBuilder.BuildWithStringPayload(stringPayload)
	if err != nil {
		panic(err)
	}

	topic := resource.TopicOf(topicId)

	publishErr := directPublisher.Publish(message, topic)
	if publishErr != nil {
		panic(publishErr)
	}

	p.logger.Println("Message Topic: ", topic.GetName())

	//directPublisher.Terminate(1 * time.Second)
	//ms.Disconnect()
	p.logger.Println("\nDirect Publisher Terminated? ", directPublisher.IsTerminated())
	return runtime.ToValue(directPublisher).ToObject(runtime)
}
func (p *DirectPublish) DisconnectPublisher(dp solace.DirectMessagePublisher) {
	dp.Terminate(0)
}
