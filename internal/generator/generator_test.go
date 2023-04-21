package generator

import (
	"github.com/lucasloureiror/AegisPass/internal/config"
	"testing"
	"unicode/utf8"
)

func TestGeneratePassWithCorrectSize(t *testing.T) {
	pwd := config.Password{
		Size: 7,
	}

	Start(&pwd)
	got := utf8.RuneCountInString(pwd.Generated)

	if got != pwd.Size {
		t.Errorf("GeneratePass received %d, but returned pass with size %d", 7, got)
	}
}

func TestMakeRandomPass(t *testing.T) {
	pwd := config.Password{
		Size:    8,
		CharSet: []byte("ab"),
	}
	indexes := []string{"0", "0", "0", "0", "1", "1", "1", "1"}

	want := "aaaabbbb"

	makeRandomPass(&pwd, indexes)

	got := pwd.Generated

	if got == want {
		t.Errorf("Expected different result, sent %s and got %s", got, want)
	}

}
