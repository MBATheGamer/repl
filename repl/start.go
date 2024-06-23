package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/MBATheGamer/lang_core/lexer"
	"github.com/MBATheGamer/lang_core/parser"
	"github.com/MBATheGamer/repl/evaluator"
)

func Start(input io.Reader, output io.Writer) {
	var scanner = bufio.NewScanner(input)

	for {
		fmt.Print(PROMPT)
		var scanned = scanner.Scan()

		if !scanned {
			return
		}

		var line = scanner.Text()

		if line == "exit" || line == ".exit" {
			return
		}

		var lexer = lexer.New(line)

		var parser = parser.New(lexer)

		var program = parser.ParseProgram()

		if len(parser.Errors()) != 0 {
			printParserErrors(output, parser.Errors())
			continue
		}

		var evaluated = evaluator.Eval(program)

		if evaluated != nil {
			io.WriteString(output, evaluated.Inspect())
			io.WriteString(output, "\n")
		}
	}
}
