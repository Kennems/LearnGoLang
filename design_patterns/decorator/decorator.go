package decorator

type Component interface {
	Operation() string
}

type ConcreteComponent struct{}
func (c *ConcreteComponent) Operation() string { return "具体对象的操作" }

type Decorator struct {
	component Component
}
func (d *Decorator) SetComponent(c Component) { d.component = c }
func (d *Decorator) Operation() string {
	if d.component != nil {
		return d.component.Operation()
	}
	return ""
}

type ConcreteDecoratorA struct {
	Decorator
	addedState string
}
func (c *ConcreteDecoratorA) Operation() string {
	c.addedState = "New State"
	return c.Decorator.Operation() + " 具体装饰对象A的操作"
}

type ConcreteDecoratorB struct {
	Decorator
}
func (c *ConcreteDecoratorB) Operation() string {
	return c.Decorator.Operation() + " 具体装饰对象B的操作"
}
