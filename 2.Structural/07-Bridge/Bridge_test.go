package Bridge

import "testing"

func TestBridge(t *testing.T) {
	ab := &RefinedAbstraction{}
	ca := &ConcreteImplementorA{implementor: Implementor{abstraction: ab}}
	ca.OperationImp()

	cb := &ConcreteImplementorB{implementor: Implementor{abstraction: ab}}
	cb.OperationImp()
}
