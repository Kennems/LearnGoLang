package interpreter
import "testing"
func TestInterpreter(t *testing.T) {
	context := &Context{}
	list := []AbstractExpression{
		&TerminalExpression{},
		&NonterminalExpression{},
		&TerminalExpression{},
	}
	for _, exp := range list {
		exp.Interpret(context)
	}
	if context.Output != "终端解释器非终端解释器终端解释器" {
		t.Errorf("Unexpected output: %s", context.Output)
	}
}
