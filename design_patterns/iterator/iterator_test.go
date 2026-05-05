package iterator
import "testing"
func TestIterator(t *testing.T) {
	a := NewConcreteAggregate()
	a.Set(0, "A")
	a.Set(1, "B")
	a.Set(2, "C")

	i := a.CreateIterator()
	res := ""
	for !i.IsDone() {
		res += i.CurrentItem().(string)
		i.Next()
	}
	if res != "ABC" {
		t.Errorf("Unexpected result: %s", res)
	}
}
