package Bridge

import "fmt"

// Abstraction
type Abstraction interface {
	Operation()
}

// RefinedAbstraction
type RefinedAbstraction struct{}

func (r *RefinedAbstraction) Operation() {
	fmt.Println("RefinedAbstraction")
}

// Implementor
type Implementor struct {
	abstraction Abstraction
}

// ConcreteImplementorA
type ConcreteImplementorA struct {
	implementor Implementor
}

func (ca *ConcreteImplementorA) OperationImp() {
	ca.implementor.abstraction.Operation()
	fmt.Println("ConcreteImplementorA")

}

// ConcreteImplementorB
type ConcreteImplementorB struct {
	implementor Implementor
}

func (cb *ConcreteImplementorB) OperationImp() {
	cb.implementor.abstraction.Operation()
	fmt.Println("ConcreteImplementorB")

}
