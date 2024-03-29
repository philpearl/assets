package lib

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
)

// CDN is for serving static files via a CDN.
//
// It provides a template function to map a static file name to a full name including CDN prefix
// and version numbering
// It helps you create versioned .css and .js files.
type CDN struct {
	// Prefix is pre-prended to the file name, so files can be loaded via a CDN or similar
	Prefix string
	// Versions maps static file names without version information to the real name.
	Versions map[string]string
}

func NewCDN(prefix string) *CDN {
	return &CDN{
		Prefix:   prefix,
		Versions: make(map[string]string),
	}
}

// StaticFileName is intended for use as a template function to convert filenames like "js/myapp.js" to
// "https://mycdn.com/path/js/myapp.3472abcd12.js". It adds versioning information only if there is a
// versioned copy of the current file available
func (cdn *CDN) StaticFileName(name string) string {
	realname, ok := cdn.Versions[name]
	if ok {
		return cdn.Prefix + realname
	}
	return cdn.Prefix + name
}

// FindVersionedFiles finds the version numbers for current files and records them.
func (cdn *CDN) FindVersionedFiles(filesys http.FileSystem, paths []string) error {
	// Look for non-versioned files in the list of paths. Work out what the current version
	// should be named and check it exists
	for _, path := range paths {
		ext := filepath.Ext(path)
		if !isInterestingExt(ext) {
			continue // not the kind of file we care about
		}

		// Is this already a versioned file?
		dir, file := filepath.Split(path)
		filebase := file[:len(file)-len(ext)]
		if versionedFileBase.MatchString(filebase) {
			continue // is already versioned
		}

		// We need to read the content to work out the version number
		f, err := filesys.Open(path)
		if err != nil {
			return fmt.Errorf("failed to open %s. %w", path, err)
		}

		fi, err := f.Stat()
		if err != nil {
			return fmt.Errorf("failed to stat %s, %w", path, err)
		}
		data := make([]byte, fi.Size())
		n, err := f.Read(data)
		if err != nil {
			return fmt.Errorf("failed to read %s, %w", path, err)
		}
		if n != len(data) {
			return fmt.Errorf("did not read all of %s. Got %d bytes out of %d", path, n, len(data))
		}

		version := VersionNumberForFile(data)
		newName := dir + versionedName(filebase, ext, version)

		// Does the versioned file exist in the filesystem?
		_, err = filesys.Open(newName)
		if err != nil {
			continue // Versioned file isn't present
		}
		cdn.Versions[path] = newName
		log.Printf("assets: %s -> %s", path, newName)
	}
	return nil
}
