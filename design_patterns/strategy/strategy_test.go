package strategy
import "testing"
func TestStrategy(t *testing.T) {
	context := NewContext(&ConcreteStrategyA{})
	if res := context.ContextInterface(); res != "算法A实现" { t.Errorf("Error: %s", res) }

	context = NewContext(&ConcreteStrategyB{})
	if res := context.ContextInterface(); res != "算法B实现" { t.Errorf("Error: %s", res) }
}
