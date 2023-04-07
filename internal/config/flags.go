package config

import (
	"flag"
	"strings"
)

func ParseFlags(pwd *Password) {
	flag.String("numeric", " ", "Password with numbers only")

	flag.Parse()

	for i := range flag.Args() {
		if strings.Contains(flag.Arg(i), "numeric") {
			pwd.UseOnlyNums = true
			return
		}

		if strings.Contains(flag.Arg(i), "standard") {
			pwd.UseStandard = true
			return
		}

	}
}
