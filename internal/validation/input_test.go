package validation

import (
	"testing"
)

func TestGenerateInvalidPassSizes(t *testing.T) {
	size := -1
	got := sizeCheck(size)

	if got == nil {
		t.Errorf("Expected error, but received none with password size %d", size)
	}

	size = 3

	got = sizeCheck(size)

	if got == nil {
		t.Errorf("Expected error, but received none with password size %d", size)
	}

	size = 26

	got = sizeCheck(size)

	if got == nil {
		t.Errorf("Expected error, but received none with password size %d", size)
	}
}

func TestFetchSize(t *testing.T) {

	args := []string{"path", "teste"}

	size, got := fetchSize(&args)

	if got == nil {
		t.Errorf("Expected error, but received none with args slice %s and returned %d size", args, size)
	}

}
