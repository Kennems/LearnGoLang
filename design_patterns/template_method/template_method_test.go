package template_method
import "testing"
func TestTemplateMethod(t *testing.T) {
	c := &ConcreteClassA{}
	tmpl := NewTemplate(c)
	if res := tmpl.TemplateMethod(); res != "具体类A方法1实现 具体类A方法2实现" {
		t.Errorf("Error: %s", res)
	}

	c2 := &ConcreteClassB{}
	tmpl2 := NewTemplate(c2)
	if res := tmpl2.TemplateMethod(); res != "具体类B方法1实现 具体类B方法2实现" {
		t.Errorf("Error: %s", res)
	}
}
