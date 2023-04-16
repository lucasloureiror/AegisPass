package charsets

import "github.com/lucasloureiror/AegisPass/internal/config"

func Create(pwd *config.Password) {

	var (
		nums = []byte("0123456789")
		full = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!@#$%&*")
	)

	if pwd.Flags.UseOnlyNums {
		pwd.CharSet = append(pwd.CharSet, nums...)
		return
	} else {
		pwd.CharSet = full

	}

}
