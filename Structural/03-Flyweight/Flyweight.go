package Flyweight

import "fmt"

// Flyweight
type Flyweight interface {
	Operation(string)
	GetState()
}

// ConcreateFlyweight
type ConcreateFlyweight struct {
	extrinsicState string
}

func (c *ConcreateFlyweight) Operation(extrinsicState string) {
	c.extrinsicState = extrinsicState
}

func (c *ConcreateFlyweight) GetState() {
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
		fw = new(ConcreateFlyweight)
		fw.Operation(key)
		ff.pool[key] = fw
	}
	return fw
}
func (ff *FlyweightFactory) GetFlyweightCount() {
	fmt.Println(len(ff.pool))
}
