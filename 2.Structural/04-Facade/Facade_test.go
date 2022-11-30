package Facade

import "testing"

func TestFacade(t *testing.T) {
	facade := &Facade{
		SubsystemA: &SubsystemA{name: "SubsystemAa"},
		SubsystemB: &SubsystemB{name: "SubsystemBb"},
	}
	facade.Operation()
}
