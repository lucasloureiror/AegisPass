package randomclient

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/lucasloureiror/AegisPass/internal/config"
)

func Init(pwd *config.Password) []string {

	client := &http.Client{}

	response, _ := fetchAPI(client, pwd)

	return response
}

type HTTPClient interface {
	Get(url string) (resp *http.Response, err error)
}

func fetchAPI(client HTTPClient, pwd *config.Password) ([]string, error) {

	url := fmt.Sprintf("https://www.random.org/integers/?num=%d&min=0&max=%d&col=1&base=10&format=plain&rnd=new", pwd.Size, len(pwd.CharSet)-1)

	response, responseError := client.Get(url)

	if responseError != nil {
		fmt.Println("Error while fetching random.org API")
		return nil, responseError
	}

	responseByte, _ := io.ReadAll(response.Body)

	responseArr := strings.Split(string(responseByte), "\n")

	return responseArr, nil
}
