package shuffle

import (
	"bytes"
	"testing"
)

func TestShuffle(t *testing.T) {

	want := []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!@#$%&*")

	got := Shuffle()
	if bytes.Equal(got, want){
		t.Error("Array wasn't shuffled")
	}

}