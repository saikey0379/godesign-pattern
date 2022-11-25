package Adapter

import "testing"

func TestAdapter(t *testing.T) {
	target := Rmvb{}
	target.Play()

	adapter := Adapter{}
	adapter.Play()
}
