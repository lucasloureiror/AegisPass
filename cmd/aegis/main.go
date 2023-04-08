package main

import (
	"github.com/lucasloureiror/AegisPass/internal/config"
	"github.com/lucasloureiror/AegisPass/internal/generator"
	"github.com/lucasloureiror/AegisPass/internal/validation"
)

func main() {
	var pwd config.Password

	config.ParseFlags(&pwd.Flags)

	validation.Init(&pwd)

	generator.Init(&pwd)

}
