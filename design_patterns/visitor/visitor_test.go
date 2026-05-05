package visitor
import "testing"
func TestVisitor(t *testing.T) {
	o := &ObjectStructure{}
	o.Attach(&ConcreteElementA{})
	o.Attach(&ConcreteElementB{})

	v := &ConcreteVisitor1{}
	res := o.Accept(v)
	expected := "访问者1 访问了 具体元素A的操作\n访问者1 访问了 具体元素B的操作\n"
	if res != expected {
		t.Errorf("Expected:\n%s\nGot:\n%s", expected, res)
	}
}
