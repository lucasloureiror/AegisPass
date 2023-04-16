package config

import (
	"flag"
	"os"
	"strings"
)

func ParseFlags(flags *Flags) {
	flag.Bool("numeric", false, "Password with numbers only")
	flag.Bool("credits", false, "Print random.org API credits to the user")
	flag.Bool("help", false, "Help the user to use the CLI tool")
	flag.Bool("standard", false, "Generate password with one upper case, one number and one special character at least.")

	flag.Parse()

	for i := range os.Args {

		if strings.Contains(os.Args[i], "help") {
			flags.NeedHelp = true
			return
		}

		if strings.Contains(os.Args[i], "credits") {
			flags.PrintCredits = true
			continue
		}

		if strings.Contains(os.Args[i], "numeric") {
			flags.UseOnlyNums = true
			return
		}

		if strings.Contains(os.Args[i], "standard") {
			flags.UseStandard = true
			return
		}

	}

}
