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
	"bytes"
	"github.com/lucasloureiror/AegisPass/internal/cli"
	"testing"
)

func TestShuffleByte(t *testing.T) {

	pwd := cli.Input{
		CharSet: []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!@#$%&*"),
	}

	want := []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!@#$%&*")

	Byte(&pwd.CharSet)

	got := pwd.CharSet
	if bytes.Equal(got, want) {
		t.Error("Array wasn't shuffled")
	}

}
