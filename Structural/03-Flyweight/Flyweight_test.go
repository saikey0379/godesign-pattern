package Flyweight

import (
	"testing"
)

func TestFlyweight(t *testing.T) {
	ff := NewFlyweightFactory()
	ff.GetFlyweightCount()

	ff.GetFlyweight("a").GetState()
	ff.GetFlyweightCount()

	ff.GetFlyweight("b").GetState()
	ff.GetFlyweightCount()
}
