package Strategy

import "testing"

func TestStrategy(t *testing.T) {
	c := &Context{}
	
	c.strategy = &ConcreteStrategyA{}
	c.ContextInterface()

	c.strategy = &ConcreteStrategyB{}
	c.ContextInterface()

	c.strategy = &ConcreteStrategyC{}
	c.ContextInterface()

}
