package main

import (
	"fmt"
	"golang.design/x/clipboard"
	"os"
	"strconv"
	"./../internal/generator"
)

func main() {
	var password_size int
	var err_convert error

	if len(os.Args) < 2 || os.Args[1] == "" {
		fmt.Println("Enter password size")
		fmt.Scan(&password_size)
	} else {
		password_size, _ = strconv.Atoi(os.Args[1])

		if err_convert != nil {
			panic("Not able to convert OS Arg to int")
		}

	password := "teste"

	fmt.Println(password)

	clipboard.Write(clipboard.FmtText, []byte(password))
}
}