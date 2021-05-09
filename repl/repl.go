package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/salleaffaire/gorilla/evaluator"
	"github.com/salleaffaire/gorilla/lexer"
	"github.com/salleaffaire/gorilla/object"
	"github.com/salleaffaire/gorilla/parser"
)

const VERSION = "0.0.1"

const PROMPT = ">>"

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, "Parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}

func Start(in io.Reader, out io.Writer) {

	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()
	ctx := object.NewYieldContext()

	fmt.Printf("Gorilla %s REPL\n", VERSION)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()

		if !scanned {
			return
		}

		line := scanner.Text()

		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()

		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env, ctx)

		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}

	}
}
