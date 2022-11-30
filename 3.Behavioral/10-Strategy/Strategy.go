package Strategy

import "fmt"

// Strategy
type Strategy interface {
	AlgorithmInterface()
}

// ConcreteStrategyA
type ConcreteStrategyA struct{}

func (cs *ConcreteStrategyA) AlgorithmInterface() {
	fmt.Println("ConcreteStrategyA")
}

// ConcreteStrategyB
type ConcreteStrategyB struct{}

func (cs *ConcreteStrategyB) AlgorithmInterface() {
	fmt.Println("ConcreteStrategyB")
}

// ConcreteStrategyC
type ConcreteStrategyC struct{}

func (cs *ConcreteStrategyC) AlgorithmInterface() {
	fmt.Println("ConcreteStrategyC")

}

// Context
type Context struct {
	strategy Strategy
}

func (c *Context) ContextInterface() {
	c.strategy.AlgorithmInterface()
}
