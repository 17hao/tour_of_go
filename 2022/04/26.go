package main

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"shiqihao.xyz/tour-of-go/antlr/calc/parser"
)

type MyListener struct {
	*parser.BaseCalcListener
}

func NewMyListener() *MyListener {
	return new(MyListener)
}

func (l *MyListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println(ctx.GetText())
}

func main() {
	lexer := parser.NewCalcLexer(antlr.NewInputStream("(1+3)*5"))
	p := parser.NewCalcParser(antlr.NewCommonTokenStream(lexer, 0))
	tree := p.Prog()
	listener := NewMyListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
}

// ===============
func Foo26() int {
	return 0x7fffffff
}
