package main

import (
	"fmt"
	"github.com/lucasloureiror/AegisPass/internal/config"
	"github.com/lucasloureiror/AegisPass/internal/generator"
	"os"
	"strconv"
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
			panic("Not able to convert OS Arg to int")
		}
	}

	config.ParseFlags()

	generator.Init(&pwd)

	fmt.Println(pwd.Generated)
}
