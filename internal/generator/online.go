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
	"strconv"
	"sync"

	"github.com/lucasloureiror/AegisPass/internal/cli"
	"github.com/lucasloureiror/AegisPass/internal/randomclient"
	"github.com/lucasloureiror/AegisPass/internal/shuffle"
)

type online struct{}

func (online) generate(input *cli.Input) (string, int, error) {

	var randomIndex []string
	var wg sync.WaitGroup
	var credits int

	wg.Add(2)

	go func() {
		defer wg.Done()
		randomIndex, credits, _ = randomclient.Start(input)
	}()

	go func() {
		defer wg.Done()
		shuffle.Byte(&input.CharSet)
	}()

	var index int
	var generated string
	pwdLen := input.Size - len(generated)

	wg.Wait()

	for i := 0; i < pwdLen; i++ {
		index, _ = strconv.Atoi(randomIndex[i])
		generated = generated + string(input.CharSet[index])
	}
	shuffle.String(&generated)

	return generated, credits, nil

}
