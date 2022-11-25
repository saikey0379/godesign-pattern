package AbstractFactory

import "fmt"

const (
	ManufacturerAmd = "AMD"
	ManufacturerItl = "Intel"
)

// AbstractProductA
type Cpu interface {
	GetProductInfo()
}

// ProductA1
type CpuAmd struct{}

func (p *CpuAmd) GetProductInfo() {
	fmt.Println(ManufacturerAmd, "central processing unit")
}

// ProductA2
type CpuItl struct{}

func (p *CpuItl) GetProductInfo() {
	fmt.Println(ManufacturerItl, "central processing unit")
}

// AbstractProductB
type Gpu interface {
	GetProductInfo()
}

// ProductB1
type GpuAmd struct{}

func (p *GpuAmd) GetProductInfo() {
	fmt.Println(ManufacturerAmd, "graphics processing unit")
}

// ProductB2
type GpuItl struct{}

func (p *GpuItl) GetProductInfo() {
	fmt.Println(ManufacturerItl, "graphics processing unit")
}

// AbstractFactory
type IAbstractFactory interface {
	CreateProductCpu(manufacturer string) Cpu
	CreateProductGpu(manufacturer string) Gpu
}

// ConcreateFactory1
type ConcreateFactory1 struct{}

// CreateProductA
func (c *ConcreateFactory1) CreateProductCpu(manufacturer string) Cpu {
	switch manufacturer {
	case "AMD":
		return &CpuAmd{}
	case "Intel":
		return &CpuItl{}
	}
	return nil
}

// CreateProductB
func (c *ConcreateFactory1) CreateProductGpu(manufacturer string) Gpu {
	return nil
}

// ConcreateFactory2
type ConcreateFactory2 struct{}

// CreateProductA
func (g *ConcreateFactory2) CreateProductCpu(manufacturer string) Cpu {
	return nil
}

// CreateProductB
func (g *ConcreateFactory2) CreateProductGpu(manufacturer string) Gpu {
	switch manufacturer {
	case "AMD":
		return &GpuAmd{}
	case "Intel":
		return &GpuItl{}
	}
	return nil
}

// Factory
func NewFactory(factory string) IAbstractFactory {
	switch factory {
	case "CPU":
		return &ConcreateFactory1{}
	case "GPU":
		return &ConcreateFactory2{}
	default:
		return nil
	}
}
