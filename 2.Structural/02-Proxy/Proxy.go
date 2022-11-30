package Proxy

import "fmt"

// Subject
type Subject interface {
	Request()
}

// RealSubject
type RealSubject struct{}

func (r *RealSubject) Request() {
	fmt.Println("Request...")
}

// Proxy
type Proxy struct {
	subject *RealSubject
}

func NewSubject() *Proxy {
	return &Proxy{subject: &RealSubject{}}
}
func (p *Proxy) Request() {
	p.subject.Request()
}
