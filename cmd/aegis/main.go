package main

import (
	"fmt"
	"github.com/lucasloureiror/AegisPass/internal/generator"
	"os"
	"strconv"
)

func main() {
	var password_size int
	var err_convert error

	if len(os.Args) < 2 || os.Args[1] == "" {
		fmt.Println("Enter password size")
		fmt.Scan(&password_size)
	} else {
		password_size, err_convert = strconv.Atoi(os.Args[1])

		if err_convert != nil {
			panic("Not able to convert OS Arg to int")
		}
	}

	password := generator.GeneratePass(password_size)

	fmt.Print(password)
}
