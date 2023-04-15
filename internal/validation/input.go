package validation

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/lucasloureiror/AegisPass/internal/config"
)

func Init(pwd *config.Password) {

	var fetchErr error

	pwd.Size, fetchErr = fetchSize(&os.Args)
	if fetchErr != nil {
		os.Exit(2)
	}

	sizeError := sizeCheck(pwd.Size)
	if sizeError != nil {
		os.Exit(2)
	}

	flags(&pwd.Flags)

}

func sizeCheck(size int) error {

	if size <= 3 || size > 16 {
		error := errors.New("password size must be bigger than 3 and smaller than 16")
		fmt.Println("Error:", error.Error())
		return error

	}
	return nil
}

func fetchSize(args *[]string) (int, error) {

	var convertErr error
	var size int

	if len(*args) < 2 || (*args)[1] == "" {
		fmt.Println("Enter password size")
		fmt.Scan(&size)
	} else {
		size, convertErr = strconv.Atoi((*args)[1])
	}

	if convertErr != nil {
		convertErr = errors.New("not able to convert OS Arg to int, did you put the number on first argument?")
		fmt.Println("Error:", convertErr.Error())
		return size, convertErr
	}

	return size, nil

}

func flags(flags *config.Flags) {

}
