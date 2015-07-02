package test

// This is from github.com/elazarl/go-bindata-assetfs
//go:generate go-bindata-assetfs -prefix "data/" -pkg "test" data/...

import (
	"net/http"
)

var FileSystem http.FileSystem = assetFS()
var Paths []string = AssetNames()

var Statics http.Handler = http.StripPrefix("/static/", http.FileServer(assetFS()))
