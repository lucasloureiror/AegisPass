package randomclient

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func FetchAPI(pwdSize int) []string {

	url := fmt.Sprintf("https://www.random.org/integers/?num=%d&min=0&max=68&col=1&base=10&format=plain&rnd=new", pwdSize)
	response, responseError := http.Get(url)

	if responseError != nil {
		panic("Error while fetching random.org API")
	}

	responseByte, _ := io.ReadAll(response.Body)

	responseArr := strings.Split(string(responseByte), "\n")

	return responseArr
}
