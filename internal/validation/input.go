package validation

import (
	"errors"
	"github.com/lucasloureiror/AegisPass/internal/config"
	"github.com/lucasloureiror/AegisPass/internal/output"
	"os"
	"strconv"
)

func Init(pwd *config.Password) {
	var fetchErr error
	flags(&pwd.Flags)

	pwd.Size, fetchErr = fetchSize(&os.Args)
	if fetchErr != nil {
		os.Exit(2)
	}

	sizeError := sizeCheck(pwd.Size)
	if sizeError != nil {
		os.Exit(2)
	}

}

func sizeCheck(size int) error {

	if size <= 3 || size > 16 {
		err := errors.New("password size must be bigger than 3 and smaller than 16")
		output.PrintError(err.Error())
		return err

	}
	return nil
}

func fetchSize(args *[]string) (int, error) {

	var convertErr error
	var size int

	if len(*args) < 2 || (*args)[1] == "" {
		size = 10
	} else {
		size, convertErr = strconv.Atoi((*args)[1])
	}

	if convertErr != nil {
		convertErr = errors.New("not able to convert OS Arg to int, did you put the number on first argument?")
		output.PrintError(convertErr.Error())
		return size, convertErr
	}

	return size, nil

}

func flags(flags *config.Flags) error {

	if flags.NeedHelp {
		output.PrintHelp()
		return nil
	}

	return nil
}
