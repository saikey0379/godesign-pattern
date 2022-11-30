package Observer

import "testing"

func TestObserver(t *testing.T) {
	ob1 := &ConcreteObserver{Name: "observer1"}
	ob2 := &ConcreteObserver{Name: "observer2"}

	cs := &ConcreteSubject{}
	cs.event = "Heavy snow today, temperature -5 degrees to 0 degrees"
	cs.Attach(ob1)
	cs.Notify()

	cs.Attach(ob2)
	cs.Notify()

	cs.Detach(ob1)
	cs.Notify()
}
