package Interpreter

import "testing"

func TestInterpreter(t *testing.T) {
	terminal := Context{"terminal"}
	tinterpreter := &TerminalExpression{value: terminal}
	tinterpreter.Interpret()

	ninterpreter := &NonterminalExpression{value: terminal}
	ninterpreter.Interpret()
}
