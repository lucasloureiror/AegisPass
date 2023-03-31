package generator_test

import (
	"testing"
	"unicode/utf8"
	"github.com/lucasloureiror/AegisPass/internal/generator"
)

func TestGeneratePassSize(t *testing.T){
	size := 7
	got := utf8.RuneCountInString(generator.GeneratePass(7))

	if got != size + 1{
		t.Errorf("GeneratePass received %d, but returned pass with size %d", 7, got)
	}
}

func TestGeneratePassSizeValidate(t *testing.T){
	size := 7
	got := utf8.RuneCountInString(generator.GeneratePass(7))

	if got != size + 1{
		t.Errorf("GeneratePass received %d, but returned pass with size %d", 7, got)
	}
}