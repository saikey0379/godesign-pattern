package FactoryMethod

import "fmt"

const (
	ManufacturerAmd = "AMD"
	ManufacturerItl = "Intel"
)

// Product
type IProduct interface {
	GetManufacturer()
}

// Creator
type IFactory interface {
	NewProduct() IProduct
}

// ConcreteProduct
type CpuAmd struct{}

func (p *CpuAmd) GetManufacturer() {
	fmt.Println(ManufacturerAmd)
}

type CpuItl struct{}

func (p *CpuItl) GetManufacturer() {
	fmt.Println(ManufacturerItl)
}

// ConcreteCreator
type FactoryAmd struct{}

func (*FactoryAmd) NewProduct() IProduct {
	return &CpuAmd{}
}

type FactoryItl struct{}

func (*FactoryItl) NewProduct() IProduct {
	return &CpuItl{}
}
