package decorator
import "testing"
func TestDecorator(t *testing.T) {
	c := &ConcreteComponent{}
	d1 := &ConcreteDecoratorA{}
	d2 := &ConcreteDecoratorB{}

	d1.SetComponent(c)
	d2.SetComponent(d1)

	res := d2.Operation()
	expected := "具体对象的操作 具体装饰对象A的操作 具体装饰对象B的操作"
	if res != expected {
		t.Errorf("Expected: %s, Got: %s", expected, res)
	}
}
