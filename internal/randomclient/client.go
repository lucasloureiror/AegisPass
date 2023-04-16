package randomclient

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/lucasloureiror/AegisPass/internal/config"
)

func Init(pwd *config.Password) []string {

	client := &http.Client{}

	error := fetchAPICredits(pwd)

	if error != nil {
		os.Exit(1)
	}
	response, error := fetchAPI(client, pwd)

	if error != nil {
		os.Exit(1)
	}

	return response
}

type HTTPClient interface {
	Get(url string) (resp *http.Response, err error)
}

func fetchAPI(client HTTPClient, pwd *config.Password) ([]string, error) {
	size := pwd.Size

	if pwd.Flags.UseStandard {
		size = size - 2
	}

	url := fmt.Sprintf("https://www.random.org/integers/?num=%d&min=0&max=%d&col=1&base=10&format=plain&rnd=new", size, len(pwd.CharSet)-1)

	response, responseError := client.Get(url)

	if responseError != nil {
		fmt.Println("Error while fetching random.org API")

		if netErr, ok := responseError.(net.Error); ok && netErr.Timeout() {
			fmt.Println("The server seems to be offline.")
		}
		return nil, responseError
	}

	responseBytes, _ := io.ReadAll(response.Body)

	responseArr := strings.Split(string(responseBytes), "\n")

	return responseArr, nil
}

func fetchAPICredits(pwd *config.Password) error {

	client := &http.Client{}

	url := "https://www.random.org/quota/?format=plain"

	response, responseError := client.Get(url)

	if responseError != nil {
		fmt.Println("Error while fetching random.org API")

		if netErr, ok := responseError.(net.Error); ok && netErr.Timeout() {
			fmt.Println("The server seems to be offline.")
		}
		return responseError
	}

	responseBytes, _ := io.ReadAll(response.Body)

	responseStr := strings.Split(string(responseBytes), "\n")

	credits, _ := strconv.Atoi(responseStr[0])
	pwd.APICredit = credits
	checkEnoughCredits(pwd.APICredit)

	return nil
}

func checkEnoughCredits(credits int) {

	if credits <= 200 {
		fmt.Println("Warning! You have very low credits with random.org API, consult docs for more details!")
		return
	}
	if credits < 0 {
		panic("Credits below zero! Can't proceed!")
	}

}
