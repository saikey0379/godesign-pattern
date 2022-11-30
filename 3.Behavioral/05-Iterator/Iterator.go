package Iterator

// Iterator
type Iterator interface {
	First() interface{}
	Next() interface{}
	IsDone() bool
	CurrentItem() interface{}
}

// ConcreteIterator
type ConcreteIterator struct {
	List  []interface{}
	Index int
}

func (ci *ConcreteIterator) First() interface{} {
	return ci.List[0]
}
func (ci *ConcreteIterator) Next() interface{} {
	if ci.IsDone() {
		ci.Index = ci.Index + 1
		return ci.List[ci.Index]
	}
	return nil
}
func (ci *ConcreteIterator) IsDone() bool {
	if ci.Index < len(ci.List)-1 {
		return true
	}
	return false
}
func (ci *ConcreteIterator) CurrentItem() interface{} {
	return ci.List[ci.Index]
}

// Aggregate
type Aggregate interface {
	CreateIterator()
}

// ConcreteAggregate
type ConcreteAggregate struct {
}

func (ca *ConcreteAggregate) CreateIterator(list []interface{}) *ConcreteIterator {
	return &ConcreteIterator{
		List:  list,
		Index: 0,
	}
}
