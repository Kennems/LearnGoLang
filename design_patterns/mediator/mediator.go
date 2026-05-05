package mediator

type Mediator interface {
	Send(message string, colleague Colleague) string
}

type Colleague interface {
	Send(message string) string
	Notify(message string) string
}

type ConcreteMediator struct {
	Colleague1 Colleague
	Colleague2 Colleague
}
func (m *ConcreteMediator) Send(message string, colleague Colleague) string {
	if colleague == m.Colleague1 {
		return m.Colleague2.Notify(message)
	}
	return m.Colleague1.Notify(message)
}

type ConcreteColleague1 struct {
	mediator Mediator
}
func NewConcreteColleague1(mediator Mediator) *ConcreteColleague1 { return &ConcreteColleague1{mediator: mediator} }
func (c *ConcreteColleague1) Send(message string) string { return c.mediator.Send(message, c) }
func (c *ConcreteColleague1) Notify(message string) string { return "同事1得到信息:" + message }

type ConcreteColleague2 struct {
	mediator Mediator
}
func NewConcreteColleague2(mediator Mediator) *ConcreteColleague2 { return &ConcreteColleague2{mediator: mediator} }
func (c *ConcreteColleague2) Send(message string) string { return c.mediator.Send(message, c) }
func (c *ConcreteColleague2) Notify(message string) string { return "同事2得到信息:" + message }
