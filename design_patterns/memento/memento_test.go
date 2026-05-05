package memento
import "testing"
func TestMemento(t *testing.T) {
	o := &Originator{}
	o.SetState("On")
	c := &Caretaker{}
	c.SetMemento(o.CreateMemento())

	o.SetState("Off")
	if o.State() != "Off" { t.Errorf("State should be Off") }

	o.SetMemento(c.Memento())
	if o.State() != "On" { t.Errorf("State should be On") }
}
