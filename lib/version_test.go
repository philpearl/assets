package lib

import (
	"testing"
)

func TestVersionNumberForFile(t *testing.T) {
	v := VersionNumberForFile([]byte("The quick brown fox jumps over the lazy dog."))
	if v != "e4d909c290d0fb1ca068ffaddf22cbd0" {
		t.Fatalf("Wrong version. Have %s", v)
	}
}

func TestVersionedName(t *testing.T) {
	n := versionedName("ethelred", ".unready", "e4d909c290d0fb1ca068ffaddf22cbd0")
	if n != "ethelred-e4d909c290d0fb1ca068ffaddf22cbd0.unready" {
		t.Fatalf("Wrong name. Have %s", n)
	}
}

func TestVersionRegex(t *testing.T) {
	assertMatch := func(name string, exp bool) {
		if versionedFileBase.MatchString(name) != exp {
			t.Errorf("%s match not as expected. Expected %t", name, exp)
		}
	}

	assertMatch("norman", false)
	assertMatch("d3", false)
	assertMatch("ethelred-e4d909c290d0fb1ca068ffaddf22cbd0", true)
	assertMatch("ethelred.e4d909c290d0fb1ca068ffaddf22cbd0", false)
	assertMatch("ethelred-e4d909c290d0fb1ca068ffaddf22cb", false)
	assertMatch("ethelred-e4d909c290d0FB1ca068ffaddf22cbd0", false)
}
