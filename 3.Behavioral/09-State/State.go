package State

import "fmt"

// Context
type Context struct {
	state State
}

func (c *Context) Request() bool {
	c.state.Handle()
	return true
}

// State
type State interface {
	Handle()
}

// ConcreteStateA
type ConcreteStateA struct {
}

func (cs *ConcreteStateA) Handle() {
	fmt.Println("ConcreteStateA")
}

// ConcreteStateB
type ConcreteStateB struct {
}

func (cs *ConcreteStateB) Handle() {
	fmt.Println("ConcreteStateB")
}
