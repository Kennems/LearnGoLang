package flyweight
import "strconv"

type Flyweight interface {
	Operation(extrinsicstate int) string
}

type ConcreteFlyweight struct{}
func (c *ConcreteFlyweight) Operation(extrinsicstate int) string {
	return "具体Flyweight:" + strconv.Itoa(extrinsicstate)
}

type UnsharedConcreteFlyweight struct{}
func (u *UnsharedConcreteFlyweight) Operation(extrinsicstate int) string {
	return "不共享的具体Flyweight:" + strconv.Itoa(extrinsicstate)
}

type FlyweightFactory struct {
	flyweights map[string]Flyweight
}
func NewFlyweightFactory() *FlyweightFactory {
	f := &FlyweightFactory{flyweights: make(map[string]Flyweight)}
	f.flyweights["X"] = &ConcreteFlyweight{}
	f.flyweights["Y"] = &ConcreteFlyweight{}
	f.flyweights["Z"] = &ConcreteFlyweight{}
	return f
}
func (f *FlyweightFactory) GetFlyweight(key string) Flyweight {
	return f.flyweights[key]
}
