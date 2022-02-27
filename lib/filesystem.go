package lib

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// CreateVersionedFiles walks a directory looking for unversioned .js and .css files. It creates
// versioned copies if they don't already exist.
func CreateVersionedFiles(root string) error {
	// Walk the directory tree
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		ext := filepath.Ext(path)
		if !isInterestingExt(ext) {
			return nil
		}
		// Is this already a versioned file?
		dir, file := filepath.Split(path)
		filebase := file[:len(file)-len(ext)]
		if versionedFileBase.MatchString(filebase) {
			return nil
		}

		// Lets work out the version number
		data, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("failed to read file %s. %w", path, err)
		}
		version := VersionNumberForFile(data)
		newName := filepath.Join(dir, versionedName(filebase, ext, version))

		// This file may already exist
		_, err = os.Stat(newName)
		if err == nil {
			// File already exists
			return nil
		}

		if !os.IsNotExist(err) {
			return fmt.Errorf("couldn't access %s. %w", newName, err)
		}

		// Versioned file does not exist, so we create it
		log.Printf("genversions: create %s", newName)
		if err := os.WriteFile(newName, data, info.Mode()); err != nil {
			return fmt.Errorf("failed to write file %s. %w", newName, err)
		}
		return nil
	})

	return err
}
