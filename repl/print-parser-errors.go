package repl

import "io"

func printParserErrors(output io.Writer, errors []string) {
	io.WriteString(output, "Woops! We ran into some problems here!\n")
	io.WriteString(output, " parser errors:\n")
	for _, message := range errors {
		io.WriteString(output, "\t"+message+"\n")
	}
}
