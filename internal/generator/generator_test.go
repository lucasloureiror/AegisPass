package generator

import (
	"testing"
	"unicode/utf8"
)

func TestGeneratePassWithCorrectSize(t *testing.T) {
	size := 7
	got := utf8.RuneCountInString(GeneratePass(size))

	if got != size {
		t.Errorf("GeneratePass received %d, but returned pass with size %d", 7, got)
	}
}

func TestGenerateInvalidPassSizes(t *testing.T) {
	size := -1
	got := validateSize(&size)

	if got == nil {
		t.Errorf("Expected error, but received none with password size %d", size)
	}

	size = 26

	got = validateSize(&size)

	if got == nil {
		t.Errorf("Expected error, but received none with password size %d", size)
	}
}
