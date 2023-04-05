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

	for i := 0; i < len(flag.Args()); i++ {
	
		if strings.Contains(flag.Arg(i), "numeric") {

			fmt.Println("achei a flag numeric")
		}

	}
}
