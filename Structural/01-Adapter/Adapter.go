package Adapter

import "fmt"

// Target
type ITarget interface {
	Play()
}
type Rmvb struct {
}

func (r *Rmvb) Play() {
	fmt.Println("This is a RMVB media,Start playing...")
}

// Adaptee
type Wmv struct{}

func (w *Wmv) PlayWmv() {
	fmt.Println("This is a WMV media,Start playing...")
}

// Adapter
type Adapter struct {
	*Wmv
}

func (a *Adapter) Play() {
	a.PlayWmv()
}
