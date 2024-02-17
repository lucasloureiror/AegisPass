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
package validation

import (
	"testing"
)

func TestGenerateInvalidPassSizes(t *testing.T) {
	size := -1
	got := sizeCheck(size)

	if got == nil {
		t.Errorf("Expected error, but received none with password size %d", size)
	}

	size = 3

	got = sizeCheck(size)

	if got == nil {
		t.Errorf("Expected error, but received none with password size %d", size)
	}

	size = 36

	got = sizeCheck(size)

	if got == nil {
		t.Errorf("Expected error, but received none with password size %d", size)
	}
}

func TestFetchSizeWithoutValue(t *testing.T) {

	args := "teste"
	defaultSize := 15

	size := fetchSize(args)

	if size != defaultSize {
		t.Errorf("Expected size %d, but received size %d", defaultSize, size)
	}

}
