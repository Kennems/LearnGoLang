package factory_method

import "testing"

func TestOperatorFactory(t *testing.T) {
	tests := []struct {
		name     string
		factory  OperatorFactory
		a, b     int
		expected int
	}{
		{"Plus", &PlusOperatorFactory{}, 1001, 10, 1011},
		{"Minus", &MinusOperatorFactory{}, 1001, 10, 991},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			op := tt.factory.Create()
			op.SetA(tt.a)
			op.SetB(tt.b)
			if got := op.Result(); got != tt.expected {
				t.Fatalf("expected %v, got %v", tt.expected, got)
			}
		})
	}
}
