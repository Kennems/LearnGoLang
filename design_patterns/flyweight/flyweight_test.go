package flyweight
import "testing"
func TestFlyweight(t *testing.T) {
	extrinsicstate := 22
	f := NewFlyweightFactory()
	fx := f.GetFlyweight("X")
	res := fx.Operation(extrinsicstate)
	if res != "具体Flyweight:22" {
		t.Errorf("Unexpected result: %s", res)
	}
}
