package factory_method

// Operator 接口
type Operator interface {
	SetA(int)
	SetB(int)
	Result() int
}

// 工厂接口
type OperatorFactory interface {
	Create() Operator
}

// OperatorBase 提供公共字段与方法
type OperatorBase struct {
	a, b int
}

func (o *OperatorBase) SetA(a int) { o.a = a }
func (o *OperatorBase) SetB(b int) { o.b = b }

// PlusOperator
type PlusOperator struct{ *OperatorBase }

func (p *PlusOperator) Result() int { return p.a + p.b }

// PlusOperatorFactory
type PlusOperatorFactory struct{}

func (PlusOperatorFactory) Create() Operator {
	return &PlusOperator{&OperatorBase{}}
}

// MinusOperator
type MinusOperator struct{ *OperatorBase }

func (m *MinusOperator) Result() int { return m.a - m.b }

// MinusOperatorFactory
type MinusOperatorFactory struct{}

func (MinusOperatorFactory) Create() Operator {
	return &MinusOperator{&OperatorBase{}}
}
