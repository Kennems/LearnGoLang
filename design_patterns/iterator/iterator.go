package iterator

type Iterator interface {
	First() interface{}
	Next() interface{}
	IsDone() bool
	CurrentItem() interface{}
}

type Aggregate interface {
	CreateIterator() Iterator
}

type ConcreteAggregate struct {
	items []interface{}
}
func NewConcreteAggregate() *ConcreteAggregate {
	return &ConcreteAggregate{items: make([]interface{}, 0)}
}
func (c *ConcreteAggregate) CreateIterator() Iterator {
	return NewConcreteIterator(c)
}
func (c *ConcreteAggregate) Count() int { return len(c.items) }
func (c *ConcreteAggregate) Get(index int) interface{} { return c.items[index] }
func (c *ConcreteAggregate) Set(index int, value interface{}) {
	if index >= len(c.items) {
		c.items = append(c.items, value)
	} else {
		c.items[index] = value
	}
}

type ConcreteIterator struct {
	aggregate *ConcreteAggregate
	current   int
}
func NewConcreteIterator(aggregate *ConcreteAggregate) *ConcreteIterator {
	return &ConcreteIterator{aggregate: aggregate, current: 0}
}
func (c *ConcreteIterator) First() interface{} { return c.aggregate.Get(0) }
func (c *ConcreteIterator) Next() interface{} {
	c.current++
	if c.current < c.aggregate.Count() {
		return c.aggregate.Get(c.current)
	}
	return nil
}
func (c *ConcreteIterator) IsDone() bool { return c.current >= c.aggregate.Count() }
func (c *ConcreteIterator) CurrentItem() interface{} { return c.aggregate.Get(c.current) }
