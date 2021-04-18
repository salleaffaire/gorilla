package core

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/salleaffaire/gorilla/evaluator"
	"github.com/salleaffaire/gorilla/lexer"
	"github.com/salleaffaire/gorilla/object"
	"github.com/salleaffaire/gorilla/parser"
)

func printCoreError(message string) {
	fmt.Printf("ERROR: %s\n", message)
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}

func Start(input *string, out io.Writer) object.Object {

	dat, err := ioutil.ReadFile(*input)

	if err != nil {
		printCoreError("Can't open " + *input)
		return evaluator.NULL
	}

	env := object.NewEnvironment()
	l := lexer.New(string(dat))
	p := parser.New(l)

	program := p.ParseProgram()

	if len(p.Errors()) != 0 {
		printParserErrors(out, p.Errors())
		return evaluator.NULL
	}

	evaluated := evaluator.Eval(program, env)

	return evaluated
}
