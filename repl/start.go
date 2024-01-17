package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/MBATheGamer/lang_core/lexer"
	"github.com/MBATheGamer/lang_core/parser"
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

		var lexer = lexer.New(line)

		var parser = parser.New(lexer)

		var program = parser.ParseProgram()

		if len(parser.Errors()) != 0 {
			printParserErrors(output, parser.Errors())
			continue
		}

		io.WriteString(output, program.String())
		io.WriteString(output, "\n")
	}
}
