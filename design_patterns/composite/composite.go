package composite
import "strings"

type Component interface {
	Add(c Component)
	Remove(c Component)
	Display(depth int) string
}

type Leaf struct {
	name string
}
func NewLeaf(name string) *Leaf { return &Leaf{name: name} }
func (l *Leaf) Add(c Component) {}
func (l *Leaf) Remove(c Component) {}
func (l *Leaf) Display(depth int) string {
	return strings.Repeat("-", depth) + l.name
}

type Composite struct {
	name     string
	children []Component
}
func NewComposite(name string) *Composite { return &Composite{name: name, children: make([]Component, 0)} }
func (c *Composite) Add(comp Component) { c.children = append(c.children, comp) }
func (c *Composite) Remove(comp Component) {}
func (c *Composite) Display(depth int) string {
	res := strings.Repeat("-", depth) + c.name
	for _, child := range c.children {
		res += "\n" + child.Display(depth+2)
	}
	return res
}
