package config

import (
	"flag"
	"fmt"
	"strings"
)

func ParseFlags(pwd *Password) {
	numeric := flag.String("numeric", " ", "Password with numbers only")

	flag.Parse()

	fmt.Println(flag.Args(), *numeric)

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
