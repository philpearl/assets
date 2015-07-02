package lib

import (
	"testing"
)

func TestIsInterestingExt(t *testing.T) {

	assertInteresting := func(ext string, exp bool) {
		if isInterestingExt(ext) != exp {
			t.Errorf("%s not as expected (expected %t)", ext, exp)
		}
	}

	assertInteresting(".jsss", false)
	assertInteresting(".js", true)
	assertInteresting(".css", true)
	assertInteresting(".cs", false)
	assertInteresting("js", false)
	assertInteresting("css", false)
}
