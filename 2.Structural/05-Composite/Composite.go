package Composite

import "fmt"

type Component interface {
	Add(Component)
	Remove(Component)
	GetChild(int) Component
	GetID() int
	Print()
}

type Composite struct {
	ID         int
	Name       string
	Components []Component
}

func (c *Composite) GetID() int {
	return c.ID
}

func (c *Composite) Print() {
	fmt.Println(c)
}

func (c *Composite) Add(component Component) {
	c.Components = append(c.Components, component)
}

func (c *Composite) Remove(component Component) {
	for i, v := range c.Components {
		if v == component {
			c.Components = append(c.Components[:i], c.Components[i+1:]...)
		}
	}
}

func (c *Composite) GetChild(id int) Component {
	for _, v := range c.Components {
		if v.GetID() == id {
			return v
		}
	}
	return nil
}
