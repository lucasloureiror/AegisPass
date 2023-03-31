package generator

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/lucasloureiror/AegisPass/internal/shuffle"
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

	validateSize(&passwordSize)

	charsArray:= shuffle.Shuffle()

	resultIndexArray := retrieveRandomCharsIndex(passwordSize)

	generatedPass := makeRandomPass(charsArray, resultIndexArray, passwordSize)

	return generatedPass
}

func makeRandomPass(chars[]byte, randomInt[]string, passwordSize int )(string){

	var password string
	var index int

	for i := 0; i < passwordSize; i++ {
		index,_ = strconv.Atoi(randomInt[i])
		password = password + string(chars[index])
	}


	return password

}

func validateSize(passwordSize *int){
	if *passwordSize < 1 || *passwordSize > 20{
		errors.New("Password size is not valid!")
	}
}
