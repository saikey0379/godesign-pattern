package Visitor

import "testing"

func TestVisitor(t *testing.T) {
	ceA := &ConcreteElementA{}
	ceB := &ConcreteElementB{}

	v := &ConcreteVisitor{}
	ceA.accept(v)

	ceB.accept(v)

}
