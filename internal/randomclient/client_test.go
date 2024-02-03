/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */
package randomclient

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/lucasloureiror/AegisPass/internal/config"
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
	pwd := config.Password{
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
