package test

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
)

//go:generate genversions -path=static
//go:embed static
var files embed.FS

var (
	FileSystem              = http.FS(files)
	Statics    http.Handler = http.StripPrefix("/static/", http.FileServer(FileSystem))
	Paths      []string     = assetNames()
)

func assetNames() []string {
	var paths []string

	fs.WalkDir(files, "static", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			paths = append(paths, path)
		}
		return nil
	})
	fmt.Println(paths)
	return paths
}
