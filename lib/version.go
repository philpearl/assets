package lib

import (
	"crypto/md5"
	"fmt"
	"regexp"
)

// Matches versioned filenames without an extension
var versionedFileBase = regexp.MustCompile(`-[0-9a-f]{32}$`)

func VersionNumberForFile(data []byte) string {
	chksum := md5.Sum(data)
	return fmt.Sprintf("%x", chksum)
}

func versionedName(base, ext, version string) string {
	return base + "-" + version + ext
}
