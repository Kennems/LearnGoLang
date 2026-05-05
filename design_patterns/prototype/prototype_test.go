package prototype

import "testing"

// Type1 原型类型1
type Type1 struct {
	Name string
}

func (t *Type1) Clone() Cloneable {
	tc := *t
	return &tc
}

// Type2 原型类型2
type Type2 struct {
	Name string
}

func (t *Type2) Clone() Cloneable {
	tc := *t
	return &tc
}

func setupManager() *PrototypeManager {
	manager := NewPrototypeManager()
	manager.Set("t1", &Type1{Name: "type1"})
	manager.Set("t2", &Type2{Name: "type2"})
	return manager
}

func TestPrototypeClone(t *testing.T) {
	manager := setupManager()

	tests := []struct {
		name       string
		prototype  string
		expected   string
		typeAssert string // 用于类型断言检查
	}{
		{"Type1 Clone", "t1", "type1", "Type1"},
		{"Type2 Clone", "t2", "type2", "Type2"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			obj := manager.Get(tt.prototype)
			if obj == nil {
				t.Fatalf("Get(%s) returned nil", tt.prototype)
			}

			// 验证克隆对象不是原对象
			original := manager.prototypes[tt.prototype]
			if obj == original {
				t.Fatalf("Clone returned the same object as original")
			}

			// 类型断言并检查字段
			switch tt.typeAssert {
			case "Type1":
				casted := obj.(*Type1)
				if casted.Name != tt.expected {
					t.Fatalf("expected %s, got %s", tt.expected, casted.Name)
				}
			case "Type2":
				casted := obj.(*Type2)
				if casted.Name != tt.expected {
					t.Fatalf("expected %s, got %s", tt.expected, casted.Name)
				}
			}
		})
	}
}
