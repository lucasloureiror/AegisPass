/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
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
