package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/MBATheGamer/lang_core/lexer"
	"github.com/MBATheGamer/lang_core/token"
)

const PROMPT = ">> "

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

		for tok := lexer.NextToken(); tok.Type != token.EOF; tok = lexer.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
