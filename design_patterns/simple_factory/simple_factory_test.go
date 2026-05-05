package simple_factory

import "testing"

func TestGreeting(t *testing.T) {
	tests := []struct {
		name     string
		apiType  int
		msg      string
		expected string
	}{
		{"HiType", APITypeHi, "2026", "Hi! 2026"},
		{"HelloType", APITypeHello, "2026", "Hello! 2026"},
		{"NilType", 999, "2026", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewGreeting(tt.apiType)
			if g == nil {
				if tt.expected != "" {
					t.Errorf("NewGreeting(%d) returned nil, expected non-nil", tt.apiType)
				}
				return
			}

			out := g.Say(tt.msg)
			if out != tt.expected {
				t.Errorf("Say() = %q, want %q", out, tt.expected)
			}
		})
	}
}
