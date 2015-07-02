# Assets
[![Build Status](https://travis-ci.org/philpearl/assets.svg)](https://travis-ci.org/philpearl/assets) [![GoDoc](https://godoc.org/github.com/philpearl/assets?status.svg)](https://godoc.org/github.com/philpearl/assets)


- Keep different versions of .js and .css files live on your site
- Automatically generate versioned names for static files in your templates
- Serve files via a CDN or from an alternate server

## genversions
The genversions command generates checkpointed versions of .js and .css files.  That is, if it finds a file called myapp.js, it will create a copy called myapp.abcdef1234567890.js, where the long string of hex is the MD5 checksum of the file. You can then serve multiple versions of your .js and .css files with long cache times, and if your html pages use the versioned names they will always get the version of the file they expect.

It takes a single parameter (-path) that tells it the root of the
tree it should operate on. If you do not specify a path it assumes the current directory.  Install with `go install github.com/philpearl/assets/genversions`

You can run genversions from go generate. For example this is what we have in our test package

```go
package test

//go:generate genversions -path=data
// This is from github.com/elazarl/go-bindata-assetfs
//go:generate go-bindata-assetfs -prefix "data/" -pkg "test" data/...

import (
	"net/http"
)

var FileSystem http.FileSystem = assetFS()
var Paths []string = AssetNames()

var Statics http.Handler = http.StripPrefix("/static/", http.FileServer(assetFS()))
```

## lib
Code to help you serve up versioned files as created by genversions. At the moment it pretty much assumes you are using [github.com/elazarl/go-bindata-assetfs] to bundle your static files.

```go

import (
	"html/template"
	"github.com/philpearl/assets/lib"
)

var cdn = lib.NewCDN("/") // Serving from the local server

func init() {
	cdn.FindVersionedFiles(assetFS(), AssetNames())

	funcMap := template.FuncMap{
		"static":  cdn.StaticFileName,
	}

	myTemplate, err := template.New("test").Funcs(funcmap).Parse(`<script src="{{static "static/js/a1.js"}}"></script>`)
}
```