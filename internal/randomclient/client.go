package randomclient

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func Init(pwdSize int) []string {

	client := &http.Client{}

	response, _ := fetchAPI(client, pwdSize)

	return response
}

type HTTPClient interface {
	Get(url string) (resp *http.Response, err error)
}

func fetchAPI(client HTTPClient, pwdSize int) ([]string, error) {

	url := fmt.Sprintf("https://www.random.org/integers/?num=%d&min=0&max=68&col=1&base=10&format=plain&rnd=new", pwdSize)

	response, responseError := client.Get(url)

	if responseError != nil {
		fmt.Println("Error while fetching random.org API")
		return nil, responseError
	}

	responseByte, _ := io.ReadAll(response.Body)

	responseArr := strings.Split(string(responseByte), "\n")

	return responseArr, nil
}
