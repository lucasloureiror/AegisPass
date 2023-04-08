package validation

import (
	"fmt"
	"os"
	"strconv"

	"github.com/lucasloureiror/AegisPass/internal/config"
)

func Init(pwd *config.Password) {
	size(&pwd.Size)
	flags(&pwd.Flags)

}

func size(size *int) {
	var convertErr error

	if len(os.Args) < 2 || os.Args[1] == "" {
		fmt.Println("Enter password size")
		fmt.Scan(size)
	} else {
		*size, convertErr = strconv.Atoi(os.Args[1])

		if convertErr != nil {
			panic("Not able to convert OS Arg to int, did you put the number on first argument?")
		}
	}

}

func flags(flags *config.Flags) {

}
