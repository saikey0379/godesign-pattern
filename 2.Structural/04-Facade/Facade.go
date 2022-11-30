package Facade

import "fmt"

// SubsystemA
type ISubsystemA interface {
	OperationA()
}
type SubsystemA struct {
	name string
}

func (s SubsystemA) OperationA() {
	fmt.Println(s.name)
}

// SubsystemB
type ISubsystemB interface {
	OperationB()
}
type SubsystemB struct {
	name string
}

func (s SubsystemB) OperationB() {
	fmt.Println(s.name)
}

// Facade
type IFacade interface {
	Operation()
}

type Facade struct {
	SubsystemA ISubsystemA
	SubsystemB ISubsystemB
}

func (f Facade) Operation() {
	f.SubsystemA.OperationA()
	f.SubsystemB.OperationB()
}
