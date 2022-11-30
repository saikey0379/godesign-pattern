package TemplateMethod

import "testing"

func TestTemplateMethod(t *testing.T) {
	cc := &ConcreteClass{}
	ac := new(AbstractClass)
	ac.i = cc
	ac.TemplateMethod()
}
