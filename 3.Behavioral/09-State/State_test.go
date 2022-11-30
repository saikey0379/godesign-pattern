package State

import "testing"

func TestState(t *testing.T) {
	context := &Context{}

	context.state = &ConcreteStateA{}
	if context.Request() {
		context.state = &ConcreteStateB{}
		context.Request()
	}
}
