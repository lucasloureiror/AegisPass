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
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/lucasloureiror/AegisPass/internal/cli"
)

type MockHTTPClient struct{}

func (m *MockHTTPClient) Do(*http.Request) (resp *http.Response, err error) {
	return &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(strings.NewReader("0\n1\n2\n3\n4\n5\n6\n7\n8\n9")),
	}, nil
}

func TestFetchAPI(t *testing.T) {
	mockClient := &MockHTTPClient{}
	pwd := cli.Input{
		Size: 10,
	}

	response, err := fetchAPI(mockClient, &pwd)

	if err != nil {
		t.Errorf("Expected error to be nil, but received error %s", err)
	}

	if len(response) != pwd.Size {
		t.Errorf("Incorret response from the function, was expecting %d and received %d", pwd.Size, len(response))
	}

}
