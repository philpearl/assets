package lib

import (
	"testing"
)

func TestCreateVersionedFiles(t *testing.T) {
	err := CreateVersionedFiles("../test/data/static")
	if err != nil {
		t.Fatalf("Failed: %v", err)
	}
}
