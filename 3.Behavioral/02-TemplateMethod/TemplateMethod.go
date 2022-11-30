package TemplateMethod

import "fmt"

// AbstractClass
type IAbstractClass interface {
	TemplateMethod()
	PrimitiveOperation1()
	PrimitiveOperation2()
}

type AbstractClass struct {
	i IAbstractClass
}

func (ac *AbstractClass) TemplateMethod() {
	ac.i.PrimitiveOperation1()
	ac.i.PrimitiveOperation2()
}

func (ac *AbstractClass) PrimitiveOperation1() {}

func (ac *AbstractClass) PrimitiveOperation2() {}

// ConcreteClass
type ConcreteClass struct {
}

func (cc *ConcreteClass) TemplateMethod() {
	cc.PrimitiveOperation1()
	cc.PrimitiveOperation2()
}
func (cc *ConcreteClass) PrimitiveOperation1() {
	fmt.Println("PrimitiveOperation1 ConcreteClass")
}

func (cc *ConcreteClass) PrimitiveOperation2() {
	fmt.Println("PrimitiveOperation2 ConcreteClass")
}
