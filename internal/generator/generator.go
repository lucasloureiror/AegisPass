package generator

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

func retrieveRandom(passwordSize int) string {
	url := fmt.Sprintf("https://www.random.org/strings/?num=1&len=%d&digits=on&upperalpha=on&loweralpha=on&unique=on&format=plain&rnd=new", passwordSize)
	response, responseError := http.Get(url)

	if responseError != nil {
		panic("Error while fetching random.org API")
	}

	responseByte, _ := io.ReadAll(response.Body)

	responseString := string(responseByte)

	return responseString
}

func GeneratePass(passwordSize int) string {

	validateSize(&passwordSize)

	result := retrieveRandom(passwordSize)

	return result
}

func validateSize(passwordSize *int){
	if *passwordSize < 1 || *passwordSize > 20{
		errors.New("Password size is not valid!")
	}
}
