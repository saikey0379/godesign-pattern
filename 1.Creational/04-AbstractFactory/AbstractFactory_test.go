package AbstractFactory

import "testing"

func TestNewFactory(t *testing.T) {
	cpuFactory := NewFactory("CPU")
	cpuAmd := cpuFactory.CreateProductCpu("AMD")
	cpuAmd.GetProductInfo()
	cpuIntel := cpuFactory.CreateProductCpu("Intel")
	cpuIntel.GetProductInfo()

	gpuFactory := NewFactory("GPU")
	gpuAmd := gpuFactory.CreateProductGpu("AMD")
	gpuAmd.GetProductInfo()
	gpuIntel := gpuFactory.CreateProductGpu("Intel")
	gpuIntel.GetProductInfo()
}
