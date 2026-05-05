package proxy

import "testing"

func TestProxy(t *testing.T) {
	real := &RealSubject{}
	sub := NewProxy(real)

	res := sub.Do()

	if res != "pre:real:after" {
		t.Fatalf("expect pre:real:after, got %s", res)
	}
}
