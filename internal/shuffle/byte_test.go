package shuffle

import (
	"bytes"
	"github.com/lucasloureiror/AegisPass/internal/cli"
	"testing"
)

func TestShuffleByte(t *testing.T) {

	pwd := cli.Input{
		CharSet: []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!@#$%&*"),
	}

	want := []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!@#$%&*")

	Byte(&pwd.CharSet)

	got := pwd.CharSet
	if bytes.Equal(got, want) {
		t.Error("Array wasn't shuffled")
	}

}
