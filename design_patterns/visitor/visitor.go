package visitor

type Visitor interface {
	VisitConcreteElementA(a *ConcreteElementA) string
	VisitConcreteElementB(b *ConcreteElementB) string
}

type Element interface {
	Accept(visitor Visitor) string
}

type ConcreteElementA struct{}
func (c *ConcreteElementA) Accept(visitor Visitor) string { return visitor.VisitConcreteElementA(c) }
func (c *ConcreteElementA) OperationA() string { return "具体元素A的操作" }

type ConcreteElementB struct{}
func (c *ConcreteElementB) Accept(visitor Visitor) string { return visitor.VisitConcreteElementB(c) }
func (c *ConcreteElementB) OperationB() string { return "具体元素B的操作" }

type ConcreteVisitor1 struct{}
func (c *ConcreteVisitor1) VisitConcreteElementA(a *ConcreteElementA) string {
	return "访问者1 访问了 " + a.OperationA()
}
func (c *ConcreteVisitor1) VisitConcreteElementB(b *ConcreteElementB) string {
	return "访问者1 访问了 " + b.OperationB()
}

type ObjectStructure struct {
	elements []Element
}
func (o *ObjectStructure) Attach(element Element) { o.elements = append(o.elements, element) }
func (o *ObjectStructure) Accept(visitor Visitor) string {
	res := ""
	for _, val := range o.elements {
		res += val.Accept(visitor) + "\n"
	}
	return res
}
