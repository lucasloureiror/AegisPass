package config

import (
	"flag"
	"strings"
)

func ParseFlags(flags *Flags) {
	flag.String("numeric", " ", "Password with numbers only")
	flag.String("credits", " ", "Print random.org API credits to the user")

	flag.Parse()

	for i := range flag.Args() {

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