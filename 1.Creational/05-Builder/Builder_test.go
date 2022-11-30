package Builder

import "testing"

func TestConstruct(t *testing.T) {
	director := &Director{}

	persion := &ConcreteBuilder{role: Role{parts: make([]string, 0)}}
	director.Construct(persion)
	persionResult := persion.GetResult()
	persionResult.ShowParts()
}
