package generator

import (
	"errors"
	"fmt"
	"github.com/lucasloureiror/AegisPass/internal/randomclient"
	"github.com/lucasloureiror/AegisPass/internal/shuffle"
	"os"
	"strconv"
)

func Init(pwdSize int) string {

	if (validateSize(&pwdSize)) != nil {
		os.Exit(1)
	}

	apiResponse := make(chan []string)

	go func() {
		apiResponse <- randomclient.FetchAPI(pwdSize)
	}()

	charSet := shuffle.Shuffle()

	pwd := makeRandomPass(charSet, <-apiResponse, pwdSize)

	return pwd
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
		fmt.Println(invalidSize)
		return invalidSize
	} else {
		return nil
	}
}
