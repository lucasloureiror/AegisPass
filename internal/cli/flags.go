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
package cli

import (
	"flag"
)

func ParseFlags(input *Input) {
	flag.BoolVar(&input.Flags.NeedHelp, "help", false, "Help the user to use the CLI tool")
	flag.BoolVar(&input.Flags.UseOnlyNums, "numeric", false, "Password with numbers only")
	flag.BoolVar(&input.Flags.Online, "online", false, "Generate random password using random.org API")
	flag.BoolVar(&input.Flags.PrintCredits, "credits", false, "Print random.org API credits to the user")
	flag.BoolVar(&input.Flags.UseStandard, "standard", false, "Generate password with one upper case, one number and one special character at least.")
	flag.IntVar(&input.NumberOfPasswords, "bulk", 1, "Size of the password")

	flag.Parse()
}
