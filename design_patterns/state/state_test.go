package state
import "testing"
func TestState(t *testing.T) {
	w := NewWork()
	w.SetHour(9)
	if res := w.WriteProgram(); res != "上午工作，精神百倍:9" { t.Errorf("Error: %s", res) }
	
	w.SetHour(13)
	if res := w.WriteProgram(); res != "下午状态还不错，继续努力:13" { t.Errorf("Error: %s", res) }
	
	w.SetHour(22)
	if res := w.WriteProgram(); res != "不行了，睡着了:22" { t.Errorf("Error: %s", res) }
}
