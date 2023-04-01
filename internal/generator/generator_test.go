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

func TestMakeRandomPass(t *testing.T) {
	size := 8
	chars := []byte("ab")
	indexes := []string{"0", "0", "0", "0", "1", "1", "1", "1"}

	want := "aaaabbbb"

	got := makeRandomPass(chars, indexes, size)

	if got != want {
		t.Errorf("Expected result was %s and got %s", got, want)
	}

}
