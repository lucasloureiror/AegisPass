package config

import (
	"flag"
	"strings"
)

func ParseFlags(flags *Flags) {
	flag.String("numeric", " ", "Password with numbers only")
	flag.String("credits", " ", "Print random.org API credits to the user")
	flag.String("help", " ", "Help the user to use the CLI tool")
	flag.String("standard", " ", "Generate password with one upper case, one number and one special character at least.")

	flag.Parse()

	for i := range flag.Args() {

		if strings.Contains(flag.Arg(i), "help") {
			flags.NeedHelp = true
			return
		}

		if strings.Contains(flag.Arg(i), "credits") {
			flags.PrintCredits = true
			continue
		}

		if strings.Contains(flag.Arg(i), "numeric") {
			flags.UseOnlyNums = true
			return
		}

		if strings.Contains(flag.Arg(i), "standard") {
			flags.UseStandard = true
			return
		}

	}
}
