package Command

import "fmt"

// Receiver
type IReceiver interface {
	Action()
}

type Receiver struct{}

func (r *Receiver) Action() {
	fmt.Println("Receiver Action")
}

// Command
type Command interface {
	Execute()
}

// ConcreteCommand
type ConcreteCommand struct {
	receiver IReceiver
}

func (cc *ConcreteCommand) Execute() {
	cc.receiver.Action()
}

// Invoker
type Invoker struct {
	commands []Command
}

func (i *Invoker) AddCommand(c Command) {
	i.commands = append(i.commands, c)
}

func (i *Invoker) ExecuteCommand() {
	for _, c := range i.commands {
		c.Execute()
	}
}
