package generator

import (
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

	result := retrieveRandom(passwordSize)

	return result
}
