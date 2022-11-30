package Decorator

import "fmt"

// Component
type Component interface {
	Operation()
}

// ConcreteComponent
type ConcreteComponent struct {
}

func (c *ConcreteComponent) Operation() {
	fmt.Println("ConcreteComponent")
}

// Decorator
type Decorator struct {
	c Component
}

func (d *Decorator) Operation() {
	if d.c != nil {
		d.c.Operation()
	}

}

// ConcreteDecoratorA
type ConcreteDecoratorA struct {
	Decorator
	State string
}

func (ca *ConcreteDecoratorA) Operation() {
	ca.c.Operation()
	fmt.Println(ca.State)
}

func (ca *ConcreteDecoratorA) addedState(state string) {
	ca.State = state
}

// ConcreteDecoratorB
type ConcreteDecoratorB struct {
	Decorator
}

func (cb *ConcreteDecoratorB) Operation() {
	cb.c.Operation()
	cb.AddedBehavior()
}

func (cb *ConcreteDecoratorB) AddedBehavior() {
	fmt.Println("ConcreteDecoratorB Behavior")
}
