package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/salleaffaire/gorilla/core"
	"github.com/salleaffaire/gorilla/evaluator"
	"github.com/salleaffaire/gorilla/repl"
)

func main() {

	inputFilenamePtr := flag.String("input", "", "File name of the Gorilla source program to execute")

	flag.Parse()

	if *inputFilenamePtr != "" {
		fmt.Printf("Excuting %s\n", *inputFilenamePtr)
		result := core.Start(inputFilenamePtr, os.Stdout)
		if result != nil && result != evaluator.NULL {
			io.WriteString(os.Stdout, result.Inspect())
			io.WriteString(os.Stdout, "\n")
		}
	} else {
		repl.Start(os.Stdin, os.Stdout)
	}
}
