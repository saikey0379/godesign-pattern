package Memento

type State interface{}

// Memento
type Memento struct {
	State
}

func (m *Memento) GetState() State {
	return m.State
}

func (m *Memento) SetState(state State) {
	m.State = state
}

// Originator
type Originator struct {
	State
}

func (o *Originator) SetMemento(memento Memento) {
	o.State = memento.State
}

func (o *Originator) CreateMemento() Memento {
	return Memento{o.State}
}

// Caretaker
type Caretaker struct {
	Memento
}
