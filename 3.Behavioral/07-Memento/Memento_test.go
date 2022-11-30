package Memento

import (
	"fmt"
	"testing"
)

func TestMemento(t *testing.T) {
	or := new(Originator)
	or.State = "on"

	ct := new(Caretaker)
	ct.Memento = or.CreateMemento()

	fmt.Println(ct.Memento.GetState())
	or.State = "off"
	fmt.Println(or.State)

	or.SetMemento(ct.Memento)
	fmt.Println(or.State)
}
