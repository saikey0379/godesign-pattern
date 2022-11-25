package Singleton

import (
	"sync"
)

// Singleton
var (
	product *Cpu
	once    sync.Once
)

func init() {
	once.Do(func() {
		product = &Cpu{}
	})
}

func GetProduct() *Cpu {
	return product
}

type Cpu struct {
	manufacturer string
	times        int
	mux          sync.Mutex
}

func (product *Cpu) SetManufacturer(m string) {
	product.mux.Lock()
	defer product.mux.Unlock()
	product.manufacturer = m
}
func (product *Cpu) GetManufacturer() string {
	product.mux.Lock()
	defer product.mux.Unlock()
	if product.manufacturer == "" {
		return "default"
	}
	return product.manufacturer
}

func (product *Cpu) UseProduct() {
	product.mux.Lock()
	defer product.mux.Unlock()
	product.times++
}
func (product *Cpu) GetUseTimes() int {
	product.mux.Lock()
	defer product.mux.Unlock()
	return product.times
}
