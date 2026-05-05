package mediator
import "testing"
func TestMediator(t *testing.T) {
	m := &ConcreteMediator{}
	c1 := NewConcreteColleague1(m)
	c2 := NewConcreteColleague2(m)
	m.Colleague1 = c1
	m.Colleague2 = c2

	res1 := c1.Send("吃饭了吗？")
	if res1 != "同事2得到信息:吃饭了吗？" { t.Errorf("Error: %s", res1) }
	res2 := c2.Send("没有呢，你打算请客？")
	if res2 != "同事1得到信息:没有呢，你打算请客？" { t.Errorf("Error: %s", res2) }
}
