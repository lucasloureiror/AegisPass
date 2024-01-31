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

	size = 36

	got = sizeCheck(size)

	if got == nil {
		t.Errorf("Expected error, but received none with password size %d", size)
	}
}

func TestFetchSizeWithoutValue(t *testing.T) {

	args := []string{"path", "teste"}
	defaultSize := 10

	size, _ := fetchSize(&args)

	if size != defaultSize {
		t.Errorf("Expected size %d, but received size %d", defaultSize, size)
	}

}
