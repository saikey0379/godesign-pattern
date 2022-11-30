package Mediator

import "fmt"

// Mediator
type Mediator interface {
	ForwordMessage(Colleague, string)
}

// ConcreteMediator
type ConcreteMediator struct {
	cc1 Colleague
	cc2 Colleague
}

func (cm *ConcreteMediator) ForwordMessage(colleague Colleague, message string) {
	switch colleague.(type) {
	case *ConcreteColleague1:
		cm.cc2.RecvMessage(message)
	case *ConcreteColleague2:
		cm.cc1.RecvMessage(message)
	default:
		fmt.Println("none type")
	}
}

// Colleague
type Colleague interface {
	SendMessage(message string)
	RecvMessage(message string)
}

// ConcreteColleague1
type ConcreteColleague1 struct {
	mediator Mediator
}

func (cc1 *ConcreteColleague1) SendMessage(message string) {
	cc1.mediator.ForwordMessage(cc1, message)
}

func (cc1 *ConcreteColleague1) RecvMessage(message string) {
	fmt.Println("cc1 receive: " + message)
}

// ConcreteColleague2
type ConcreteColleague2 struct {
	mediator Mediator
}

func (cc2 *ConcreteColleague2) SendMessage(message string) {
	cc2.mediator.ForwordMessage(cc2, message)
}

func (cc2 *ConcreteColleague2) RecvMessage(message string) {
	fmt.Println("cc2 receive: " + message)
}
