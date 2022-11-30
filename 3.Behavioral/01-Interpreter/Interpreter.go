package Interpreter

import "fmt"

// Context
type Context struct {
	value string
}

// AbstractExpression
type AbstractExpression interface {
	Interpret(Context)
}

// TerminalExpression
type TerminalExpression struct {
	value Context
}

func (t *TerminalExpression) Interpret() {
	fmt.Println("TerminalExpression " + t.value.value)
}

// NonterminalExpression
type NonterminalExpression struct {
	value Context
}

func (n *NonterminalExpression) Interpret() {
	fmt.Println("NonterminalExpression " + n.value.value)
}
