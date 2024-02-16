package generator

import (
	"github.com/lucasloureiror/AegisPass/internal/charsets"
	"github.com/lucasloureiror/AegisPass/internal/cli"
	"testing"
	"unicode/utf8"
)

func TestGenerateRandomPassWithCorrectSize(t *testing.T) {
	pwd := cli.Input{
		Size: 8,
	}

	random := random{}

	generator := passwordGenerator{
		data: pwd,
		mode: random,
	}

	charsets.Create(&generator.data)
	password, _, _ := generator.mode.generate(&generator.data)

	got := utf8.RuneCountInString(password)

	if got != pwd.Size {
		t.Errorf("GeneratePass received %d, but returned pass with size %d", 8, got)
	}
}
