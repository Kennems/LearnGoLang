package interpreter

type Context struct {
	Input  string
	Output string
}

type AbstractExpression interface {
	Interpret(context *Context)
}

type TerminalExpression struct{}
func (t *TerminalExpression) Interpret(context *Context) {
	context.Output += "终端解释器"
}

type NonterminalExpression struct{}
func (n *NonterminalExpression) Interpret(context *Context) {
	context.Output += "非终端解释器"
}
