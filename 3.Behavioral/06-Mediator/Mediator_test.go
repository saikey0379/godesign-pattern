package Mediator

import "testing"

func TestMediator(t *testing.T) {
	cm := &ConcreteMediator{}
	cc1 := &ConcreteColleague1{}
	cc2 := &ConcreteColleague2{}
	cc1.mediator = cm
	cc2.mediator = cm
	cm.cc1 = cc1
	cm.cc2 = cc2

	cc1.SendMessage("cc1[How are you?]")
	cc2.SendMessage("cc2[Fine thanks and you?]")
}
