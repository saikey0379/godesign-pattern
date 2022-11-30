package Decorator

import (
	"testing"
)

func TestDecorator(t *testing.T) {
	var c Component = &ConcreteComponent{}
	ca := &ConcreteDecoratorA{Decorator{c: c}, ""}
	ca.addedState("ConcreteDecoratorA start")
	ca.Operation()

	cb := &ConcreteDecoratorB{Decorator{c: ca}}
	cb.Operation()

	cc := &ConcreteDecoratorB{Decorator{c: cb}}
	cc.Operation()
}
