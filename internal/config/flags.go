package config

import (
	"flag"
	"os"
	"strings"
)

func ParseFlags(flags *Flags) {
	//Flags stay here just for better documentation of the options
	flag.Bool("numeric", false, "Password with numbers only")
	flag.Bool("credits", false, "Print random.org API credits to the user")
	flag.Bool("help", false, "Help the user to use the CLI tool")
	flag.Bool("standard", false, "Generate password with one upper case, one number and one special character at least.")
	flag.Bool("offline", false, "Generate password without using random.org API")
	args := os.Args[1:] //Removing the path

	for i := range args {

		if strings.Contains(args[i], "help") {
			flags.NeedHelp = true
			return
		} else if strings.Contains(args[i], "standard") {
			flags.UseStandard = true
			return
		}
		if strings.Contains(args[i], "credits") {
			flags.PrintCredits = true
		}

		if strings.Contains(args[i], "numeric") {
			flags.UseOnlyNums = true
		}

		if strings.Contains(args[i], "offline") {
			flags.Offline = true
		}

	}
}
