package Builder

import "fmt"

// Product
type Role struct {
	parts []string
}

func (p *Role) AddParts(s string) {
	p.parts = append(p.parts, s)
}
func (p *Role) ShowParts() {
	fmt.Println(p.parts)
}

// Builder
type Builder interface {
	Build1()
	Build2()
	Build3()
	Build4()
	GetResult() Role
}

// ConcreteBuilder
type ConcreteBuilder struct {
	role Role
}

// 增加头盔
func (c *ConcreteBuilder) Build1() {
	c.role.AddParts("Helmet")
}

// 增加铠甲
func (c *ConcreteBuilder) Build2() {
	c.role.AddParts("Armor")
}

// 增加靴子
func (c *ConcreteBuilder) Build3() {
	c.role.AddParts("Boots")
}

// 增加武器
func (c *ConcreteBuilder) Build4() {
	c.role.AddParts("Weapon")
}

func (c *ConcreteBuilder) GetResult() Role {
	return c.role
}

// Director
type Director struct{}

// Construct
func (d *Director) Construct(builder Builder) {
	builder.Build1()
	builder.Build2()
	builder.Build3()
	builder.Build4()
}
