package shuffle

import (
	"bytes"
	"testing"

	"github.com/lucasloureiror/AegisPass/internal/config"
)

func TestShuffle(t *testing.T) {

	pwd := config.Password{
		CharSet: []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!@#$%&*"),
	}

	want := []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!@#$%&*")

	ShuffleByte(&pwd.CharSet)

	got := pwd.CharSet
	if bytes.Equal(got, want) {
		t.Error("Array wasn't shuffled")
	}

}
