package command
import "testing"
func TestCommand(t *testing.T) {
	receiver := &Receiver{}
	cmd := NewConcreteCommand(receiver)
	invoker := &Invoker{}
	invoker.SetCommand(cmd)
	if res := invoker.ExecuteCommand(); res != "执行请求！" {
		t.Errorf("Unexpected result: %s", res)
	}
}
