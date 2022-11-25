package Prototype

import "fmt"

// Prototype
type IPrototype interface {
	GetManufacturer()
	Clone() IPrototype
}

// ConcreatePrototype
type CPU struct {
	Manufacturer string
}

func (p *CPU) Clone() IPrototype {
	return &CPU{Manufacturer: p.Manufacturer}
}

func (p *CPU) GetManufacturer() {
	fmt.Println(p.Manufacturer)
}
