package kata

import (
	"fmt"
	"testing"
)

func TestValidBraces(t *testing.T) {
	if err := testEmptyString(); err != nil {
		t.Fatalf("Test empty string failed: %v\n", err)
	}
	if err :=
}

func testEmptyString() error {
	isvalid := ValidBraces("")
	if isvalid == false {
		return fmt.Errorf("The empty string should return a true vallue")
	}
	return nil
}
