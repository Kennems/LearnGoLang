package adapter

import "testing"

func TestAdapter(t *testing.T) {
	adaptee := NewAdaptee()
	adapter := NewAdapter(adaptee)

	out := adapter.Request()
	expected := "specific request called"

	if out != expected {
		t.Errorf("Adapter Request() = %q, want %q", out, expected)
	}
}
