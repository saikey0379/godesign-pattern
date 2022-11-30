package Iterator

import (
	"fmt"
	"testing"
)

func TestIterator(t *testing.T) {
	var list = []interface{}{"i", "t", "e", "r", "a", "t", "o", "r"}
	ca := &ConcreteAggregate{}
	it := ca.CreateIterator(list)
	fmt.Println(it.First())
	fmt.Println(it.CurrentItem())
	fmt.Println(it.Next())
	fmt.Println(it.Next())
}
