package command

type Command interface {
	Execute() string
}

type Receiver struct{}
func (r *Receiver) Action() string { return "执行请求！" }

type ConcreteCommand struct {
	receiver *Receiver
}
func NewConcreteCommand(receiver *Receiver) *ConcreteCommand {
	return &ConcreteCommand{receiver: receiver}
}
func (c *ConcreteCommand) Execute() string {
	return c.receiver.Action()
}

type Invoker struct {
	command Command
}
func (i *Invoker) SetCommand(command Command) { i.command = command }
func (i *Invoker) ExecuteCommand() string {
	if i.command != nil { return i.command.Execute() }
	return ""
}
