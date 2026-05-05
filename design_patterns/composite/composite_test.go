package composite
import "testing"
func TestComposite(t *testing.T) {
	root := NewComposite("root")
	root.Add(NewLeaf("Leaf A"))
	comp := NewComposite("Composite X")
	comp.Add(NewLeaf("Leaf XA"))
	root.Add(comp)

	res := root.Display(1)
	expected := "-root\n---Leaf A\n---Composite X\n-----Leaf XA"
	if res != expected {
		t.Errorf("Expected:\n%s\nGot:\n%s", expected, res)
	}
}
