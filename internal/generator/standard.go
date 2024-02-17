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
package generator

import (
	"github.com/lucasloureiror/AegisPass/internal/cli"
	"github.com/lucasloureiror/AegisPass/internal/shuffle"
	"sync"
)

type standard struct{}

func (standard) generate(input *cli.Input) (string, int, error) {

	var wg sync.WaitGroup
	var credits int
	wg.Add(4)
	upper := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	special := []byte("!@#$%&*-_")
	nums := []byte("0123456789")

	go func() {
		defer wg.Done()
		shuffle.Byte(&upper)
	}()
	go func() {
		defer wg.Done()
		shuffle.Byte(&special)
	}()

	go func() {
		defer wg.Done()
		shuffle.Byte(&nums)
	}()

	go func() {
		defer wg.Done()
		shuffle.Byte(&input.CharSet)
	}()

	wg.Wait()

	requirements := string(upper[0]) + string(special[0]) + string(nums[0])

	generated := shuffle.BuildString(input.CharSet, input.Size-3)

	generated = generated + requirements
	shuffle.String(&generated)

	return generated, credits, nil
}
