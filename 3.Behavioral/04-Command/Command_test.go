package Command

import "testing"

func TestCommand(t *testing.T) {
	cc := &ConcreteCommand{&Receiver{}}
	iv := Invoker{}
	iv.AddCommand(cc)
	iv.AddCommand(cc)
	iv.ExecuteCommand()
}
