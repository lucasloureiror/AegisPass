package generator

import (
	"errors"
	"fmt"
	"github.com/lucasloureiror/AegisPass/internal/shuffle"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func retrieveRandomCharsIndex(passwordSize int) []string {

	url := fmt.Sprintf("https://www.random.org/integers/?num=%d&min=0&max=68&col=1&base=10&format=plain&rnd=new", passwordSize)
	response, responseError := http.Get(url)

	if responseError != nil {
		panic("Error while fetching random.org API")
	}

	responseByte, _ := io.ReadAll(response.Body)

	responseArray := strings.Split(string(responseByte), "\n")

	return responseArray
}

func GeneratePass(passwordSize int) string {

	if (validateSize(&passwordSize)) != nil {
		os.Exit(1)
	}

	charsArray := shuffle.Shuffle()

	resultIndexArray := retrieveRandomCharsIndex(passwordSize)

	generatedPass := makeRandomPass(charsArray, resultIndexArray, passwordSize)

	return generatedPass
}

func makeRandomPass(chars []byte, randomInt []string, passwordSize int) string {

	var password string
	var index int

	for i := 0; i < passwordSize; i++ {
		index, _ = strconv.Atoi(randomInt[i])
		password = password + string(chars[index])
	}

	return password

}

func validateSize(passwordSize *int) error {
	if *passwordSize < 1 || *passwordSize > 25 {
		invalidSize := errors.New("password size must be bigger than 0 and smaller than 25")
		fmt.Println(invalidSize)
		return invalidSize
	} else {
		return nil
	}
}
