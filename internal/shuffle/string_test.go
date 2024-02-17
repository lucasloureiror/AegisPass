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
package shuffle

import (
	"testing"
)

func TestShuffleString(t *testing.T) {

	original := "TestString1235)!@#"

	shuffled := original

	String(&shuffled)

	if original == shuffled {
		t.Error("Original string was: ", original, "Shuffled string was: ", shuffled, "Strings are equal, they shouldn't be.")
	}

}
