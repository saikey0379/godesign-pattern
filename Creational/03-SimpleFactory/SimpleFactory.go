package SimpleFactory

import "fmt"

const (
	ManufacturerAmd = "AMD"
	ManufacturerItl = "Intel"
)

type Product interface {
	GetManufacturer()
}

// product
type CpuAmd struct{}

func (p *CpuAmd) GetManufacturer() {
	fmt.Println(ManufacturerAmd)
}

type CpuItl struct{}

func (p *CpuItl) GetManufacturer() {
	fmt.Println(ManufacturerItl)
}

// factory
func SimpleFactory(product string) Product {
	switch product {
	case "AMD":
		return &CpuAmd{}
	case "Intel":
		return &CpuItl{}
	default:
		return nil
	}
}
