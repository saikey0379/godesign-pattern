package ResponsibilityChain

import "testing"

func TestResponsibilityChain(t *testing.T) {
	h1 := &ConcreteHandler1{"h1", 1, nil}
	h2 := &ConcreteHandler2{"h2", 2, h1}
	h3 := &ConcreteHandler2{"h3", 3, h2}

	h3.HandleRequest(3)
	h3.HandleRequest(2)
	h3.HandleRequest(1)

}
