package builder

import "testing"

func TestBuilder(t *testing.T) {
	tests := []struct {
		name     string
		builder  Builder
		expected interface{}
		getRes   func() interface{}
	}{
		{
			name:     "Builder1",
			builder:  &Builder1{},
			expected: "123",
			getRes: func() interface{} {
				return (&Builder1{}).GetResult()
			},
		},
		{
			name:     "Builder2",
			builder:  &Builder2{},
			expected: 6,
			getRes: func() interface{} {
				return (&Builder2{}).GetResult()
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := NewDirector(tt.builder)
			d.Construct()

			var actual interface{}
			switch b := tt.builder.(type) {
			case *Builder1:
				actual = b.GetResult()
			case *Builder2:
				actual = b.GetResult()
			default:
				t.Fatalf("unsupported builder type")
			}

			if actual != tt.expected {
				t.Fatalf("expected %v, got %v", tt.expected, actual)
			}
		})
	}
}
