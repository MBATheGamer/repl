package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/MBATheGamer/repl/repl"
)

func main() {
	var user, err = user.Current()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s!\nThis is the M programming language!\n",
		user.Username,
	)
	fmt.Println("Feel free to type in commands")
	repl.Start(os.Stdin, os.Stdout)
}
