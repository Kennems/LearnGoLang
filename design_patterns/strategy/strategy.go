package strategy

type Strategy interface {
	AlgorithmInterface() string
}

type ConcreteStrategyA struct{}
func (c *ConcreteStrategyA) AlgorithmInterface() string { return "算法A实现" }

type ConcreteStrategyB struct{}
func (c *ConcreteStrategyB) AlgorithmInterface() string { return "算法B实现" }

type ConcreteStrategyC struct{}
func (c *ConcreteStrategyC) AlgorithmInterface() string { return "算法C实现" }

type Context struct {
	strategy Strategy
}
func NewContext(strategy Strategy) *Context { return &Context{strategy: strategy} }
func (c *Context) ContextInterface() string { return c.strategy.AlgorithmInterface() }
