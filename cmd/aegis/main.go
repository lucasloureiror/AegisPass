package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/lucasloureiror/AegisPass/internal/config"
	"github.com/lucasloureiror/AegisPass/internal/generator"
)

func main() {
	var pwd config.Password
	var convertErr error

	if len(os.Args) < 2 || os.Args[1] == "" {
		fmt.Println("Enter password size")
		fmt.Scan(&pwd.Size)
	} else {
		pwd.Size, convertErr = strconv.Atoi(os.Args[1])

		if convertErr != nil {
			panic("Not able to convert OS Arg to int, did you put the number on first argument?")
		}
	}

	config.ParseFlags(&pwd)

	generator.Init(&pwd)

	fmt.Println(pwd.Generated)
}
