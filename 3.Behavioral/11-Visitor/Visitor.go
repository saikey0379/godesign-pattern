package Visitor

import "fmt"

// Visitor
type Visitor interface {
	visit(Element)
}

// ConcreteVisitor
type ConcreteVisitor struct {
}

func (cv *ConcreteVisitor) visit(element Element) {
	switch element.(type) {
	case *ConcreteElementA:
		fmt.Println("ConcreteElementA")
	case *ConcreteElementB:
		fmt.Println("ConcreteElementB")
	}
}

// Element
type Element interface {
	accept(visitor Visitor)
}

// ConcreteElementA
type ConcreteElementA struct{}

func (ce *ConcreteElementA) accept(v Visitor) {
	v.visit(ce)
}

// ConcreteElementB
type ConcreteElementB struct{}

func (ce *ConcreteElementB) accept(v Visitor) {
	v.visit(ce)
}
