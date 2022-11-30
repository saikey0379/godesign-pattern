package Flyweight

import "fmt"

// Flyweight
type Flyweight interface {
	Operation(string)
	GetState()
}

// ConcreteFlyweight
type ConcreteFlyweight struct {
	extrinsicState string
}

func (c *ConcreteFlyweight) Operation(extrinsicState string) {
	c.extrinsicState = extrinsicState
}

func (c *ConcreteFlyweight) GetState() {
	fmt.Println(c.extrinsicState)
}

// FlyweightFactory
type FlyweightFactory struct {
	pool map[string]Flyweight
}

func NewFlyweightFactory() FlyweightFactory {
	return FlyweightFactory{pool: make(map[string]Flyweight)}
}

func (ff *FlyweightFactory) GetFlyweight(key string) Flyweight {
	fw, ok := ff.pool[key]
	if !ok {
		fw = new(ConcreteFlyweight)
		fw.Operation(key)
		ff.pool[key] = fw
	}
	return fw
}
func (ff *FlyweightFactory) GetFlyweightCount() {
	fmt.Println(len(ff.pool))
}
