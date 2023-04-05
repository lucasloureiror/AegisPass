package randomclient

import (
	"io"
	"net/http"
	"strings"
	"testing"
)

type MockHTTPClient struct{}

func (m *MockHTTPClient) Get(url string) (resp *http.Response, err error) {
	return &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(strings.NewReader("0\n1\n2\n3\n4\n5\n6\n7\n8\n9")),
	}, nil
}

func TestFetchAPI(t *testing.T) {
	mockClient := &MockHTTPClient{}

	pwdSize := 10

	response, err := fetchAPI(mockClient, pwdSize)

	if err != nil {
		t.Errorf("Expected error to be nil, but received error %s", err)
	}

	if len(response) != pwdSize {
		t.Errorf("Incorret response from the function, was expecting %d and received %d", pwdSize, len(response))
	}

}
