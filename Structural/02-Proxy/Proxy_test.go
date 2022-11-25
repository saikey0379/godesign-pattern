package Proxy

import "testing"

func TestProxy(t *testing.T) {
	subject := NewSubject()
	subject.Request()
}
