package generator

import (
	"errors"
	"github.com/lucasloureiror/AegisPass/internal/config"
	"github.com/lucasloureiror/AegisPass/internal/randomclient"
	"github.com/lucasloureiror/AegisPass/internal/shuffle"
	"os"
	"strconv"
)

func Init(pwd *config.Password) {

	if (validateSize(&pwd.Size)) != nil {
		os.Exit(1)
	}

	apiResponse := make(chan []string)

	go func() {
		apiResponse <- randomclient.Init(pwd.Size)
	}()

	charSet := shuffle.Shuffle()

	pwd.Generated = makeRandomPass(charSet, <-apiResponse, pwd.Size)

}

func makeRandomPass(chars []byte, randomInt []string, pwdSize int) string {

	var pwd string
	var index int

	for i := 0; i < pwdSize; i++ {
		index, _ = strconv.Atoi(randomInt[i])
		pwd = pwd + string(chars[index])
	}

	return pwd

}

func validateSize(pwdSize *int) error {
	if *pwdSize < 1 || *pwdSize > 25 {
		invalidSize := errors.New("password size must be bigger than 0 and smaller than 25")
		return invalidSize
	} else {
		return nil
	}
}
