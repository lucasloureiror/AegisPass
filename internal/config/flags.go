package config

import (
	"flag"
	"fmt"
	"strings"
)

func ParseFlags() {
	numeric := flag.String("numeric", " ", "Password with numbers only")

	flag.Parse()

	fmt.Println(flag.Args(), *numeric)

	for i := range flag.Args() {
		if strings.Contains(flag.Arg(i), "numeric") {

			fmt.Println("achei a flag numeric")
		}

	}
}
