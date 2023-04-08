package output

import (
	"fmt"

	"github.com/lucasloureiror/AegisPass/internal/config"
)

func Print(pwd *config.Password) {
	if pwd.Flags.PrintCredits {
		fmt.Println(pwd.Generated)
		fmt.Printf("API credits remaining: %d", pwd.APICredit)
	} else {
		fmt.Println(pwd.Generated)
	}
}
