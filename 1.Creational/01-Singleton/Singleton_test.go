package Singleton

import (
	"fmt"
	"testing"
)

func TestSingletion(t *testing.T) {
	product := GetProduct()
	fmt.Println(product.GetManufacturer())
	product.SetManufacturer("Intel")
	fmt.Println(product.GetManufacturer())

	fmt.Println(product.GetUseTimes())
	product.UseProduct()
	fmt.Println(product.GetUseTimes())

}
