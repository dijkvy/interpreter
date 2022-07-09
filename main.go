package main

import (
	"fmt"
	"os"
	user2 "os/user"

	"interperter/repl"
)

func main() {
	user, err := user2.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("hello %s! This is the script programing language!\n", user.Username)
	fmt.Printf("Feel free to type in commands\n")

	repl.Start(os.Stdin, os.Stdout)

}
