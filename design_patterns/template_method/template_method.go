package template_method

type AbstractClass interface {
	PrimitiveOperation1() string
	PrimitiveOperation2() string
}

type Template struct {
	impl AbstractClass
}
func NewTemplate(impl AbstractClass) *Template { return &Template{impl: impl} }
func (t *Template) TemplateMethod() string {
	return t.impl.PrimitiveOperation1() + " " + t.impl.PrimitiveOperation2()
}

type ConcreteClassA struct{}
func (c *ConcreteClassA) PrimitiveOperation1() string { return "具体类A方法1实现" }
func (c *ConcreteClassA) PrimitiveOperation2() string { return "具体类A方法2实现" }

type ConcreteClassB struct{}
func (c *ConcreteClassB) PrimitiveOperation1() string { return "具体类B方法1实现" }
func (c *ConcreteClassB) PrimitiveOperation2() string { return "具体类B方法2实现" }
