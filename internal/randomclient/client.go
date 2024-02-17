/*
Copyright 2024 lucasloureiror

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package randomclient

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/lucasloureiror/AegisPass/internal/cli"
	"github.com/lucasloureiror/AegisPass/internal/output"
)

func Start(input *cli.Input) ([]string, int, error) {
	var err error
	var credits int

	client := &http.Client{}

	credits, err = fetchAPICredits()

	if err != nil {
		os.Exit(1)
	}
	response, error := fetchAPI(client, input)

	if error != nil {
		os.Exit(1)
	}

	return response, credits, nil
}

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

func fetchAPI(client HTTPClient, pwd *cli.Input) ([]string, error) {

	size := pwd.Size

	if pwd.Flags.UseStandard {
		size = size - 2
	}

	url := fmt.Sprintf("https://www.random.org/integers/?num=%d&min=0&max=%d&col=1&base=10&format=plain&rnd=new", size, len(pwd.CharSet)-1)

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "AegisPass github.com/lucasloureiror/AegisPass v2")

	response, responseError := client.Do(req)

	if responseError != nil {

		if response.StatusCode == http.StatusServiceUnavailable {
			output.PrintError("service unavailable")
		} else if netErr, ok := responseError.(net.Error); ok && netErr.Timeout() {
			fmt.Println("The server seems to be offline.")
		} else {
			output.PrintError("can't reach random.org api server")
		}
		return nil, responseError
	}

	defer response.Body.Close()

	responseBytes, _ := io.ReadAll(response.Body)

	responseArr := strings.Split(string(responseBytes), "\n")

	return responseArr, nil
}

func fetchAPICredits() (int, error) {
	var credits int
	client := &http.Client{}

	url := "https://www.random.org/quota/?format=plain"

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return 0, err
	}

	req.Header.Set("User-Agent", "AegisPass github.com/lucasloureiror/AegisPass")

	response, responseError := client.Do(req)

	if responseError != nil {
		fmt.Println("Error while fetching random.org API")

		if netErr, ok := responseError.(net.Error); ok && netErr.Timeout() {
			fmt.Println("The server seems to be offline.")
		}
		return 0, responseError
	}

	responseBytes, _ := io.ReadAll(response.Body)

	responseStr := strings.Split(string(responseBytes), "\n")

	credits, _ = strconv.Atoi(responseStr[0])
	checkEnoughCredits(credits)

	return credits, nil
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
