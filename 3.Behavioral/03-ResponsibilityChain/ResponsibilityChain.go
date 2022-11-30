package ResponsibilityChain

import "fmt"

// Handler
type Handler interface {
	HandleRequest(int)
}

// ConcreteHandler1
type ConcreteHandler1 struct {
	Name        string
	Level       int
	NextHandler Handler
}

func (h *ConcreteHandler1) HandleRequest(level int) {
	if h.Level == level {
		fmt.Println(h.Name)
	} else {
		h.NextHandler.HandleRequest(level)
	}
}

// ConcreteHandler2
type ConcreteHandler2 struct {
	Name        string
	Level       int
	NextHandler Handler
}

func (h *ConcreteHandler2) HandleRequest(level int) {
	if h.Level == level {
		fmt.Println(h.Level)
	} else {
		h.NextHandler.HandleRequest(level)
	}
}
