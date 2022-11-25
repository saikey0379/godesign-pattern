package Prototype

import (
	"testing"
)

func TestPrototype(t *testing.T) {
	cpuAmd := CPU{Manufacturer: "AMD"}
	cpuAmd.GetManufacturer()

	cpuAmd.Clone().GetManufacturer()
}
