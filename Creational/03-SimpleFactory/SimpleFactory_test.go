package SimpleFactory

import "testing"

func TestSimpleFactory(t *testing.T) {
	cpuAmd := SimpleFactory("AMD")
	cpuAmd.GetManufacturer()

	cpuItl := SimpleFactory("Intel")
	cpuItl.GetManufacturer()
}
