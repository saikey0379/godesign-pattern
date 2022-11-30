package Observer

import "fmt"

// Subject
type Subject interface {
	Attach(Observer)
	Detach(Observer)
	Notify()
}

// ConcreteSubject
type ConcreteSubject struct {
	event     string
	observers []Observer
}

func (cs *ConcreteSubject) Attach(observer Observer) {
	cs.observers = append(cs.observers, observer)
}

func (cs *ConcreteSubject) Detach(observer Observer) {
	if len(cs.observers) > 0 {
		for i, ob := range cs.observers {
			if ob == observer {
				cs.observers = append(cs.observers[:i], cs.observers[i+1:]...)
			}
		}
	}
}

func (cs *ConcreteSubject) Notify() {
	for _, ob := range cs.observers {
		ob.Update(cs.event)
	}
}

// Observer
type Observer interface {
	Update(string)
}

// ConcreteObserver
type ConcreteObserver struct {
	Name string
}

func (co *ConcreteObserver) Update(event string) {
	fmt.Printf("%s receive event: %v\n", co.Name, event)
}
