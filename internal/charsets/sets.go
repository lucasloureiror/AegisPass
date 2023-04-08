package charsets

import "github.com/lucasloureiror/AegisPass/internal/config"

func Create(pwd *config.Password) {

	if pwd.Flags.UseOnlyNums {
		nums := []byte("0123456789")
		pwd.CharSet = append(pwd.CharSet, nums...)
		return
	} else {
		pwd.CharSet = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!@#$%&*")

	}

}
