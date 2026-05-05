package memento

type Memento struct {
	state string
}
func NewMemento(state string) *Memento { return &Memento{state: state} }
func (m *Memento) State() string { return m.state }

type Originator struct {
	state string
}
func (o *Originator) State() string { return o.state }
func (o *Originator) SetState(state string) { o.state = state }
func (o *Originator) CreateMemento() *Memento { return NewMemento(o.state) }
func (o *Originator) SetMemento(memento *Memento) { o.state = memento.State() }

type Caretaker struct {
	memento *Memento
}
func (c *Caretaker) Memento() *Memento { return c.memento }
func (c *Caretaker) SetMemento(memento *Memento) { c.memento = memento }
