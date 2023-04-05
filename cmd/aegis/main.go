package main

import (
	"fmt"
	"github.com/lucasloureiror/AegisPass/internal/config"
	"github.com/lucasloureiror/AegisPass/internal/generator"
	"os"
	"strconv"
)

func main() {
	var pwdSize int
	var convertErr error

	if len(os.Args) < 2 || os.Args[1] == "" {
		fmt.Println("Enter password size")
		fmt.Scan(&pwdSize)
	} else {
		pwdSize, convertErr = strconv.Atoi(os.Args[1])

		if convertErr != nil {
			panic("Not able to convert OS Arg to int")
		}
	}

	config.ParseFlags()

	password := generator.Init(pwdSize)

	fmt.Println(password)
}
