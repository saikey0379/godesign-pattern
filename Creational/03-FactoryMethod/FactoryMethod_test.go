package FactoryMethod

import "testing"

func TestFactoryMethod(t *testing.T) {
	factoryAmd := &FactoryAmd{}
	factoryAmd.NewProduct().GetManufacturer()

	factoryItl := &FactoryItl{}
	factoryItl.NewProduct().GetManufacturer()
}
