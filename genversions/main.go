package main

import (
	"flag"
	"fmt"
	"github.com/philpearl/assets/lib"
	"os"
)

func main() {
	var root string

	flag.StringVar(&root, "path", ".", "Path to the root of the tree containing files to version")
	flag.Parse()

	err := lib.CreateVersionedFiles(root)
	if err != nil {
		fmt.Printf("Failed to create versioned files. %v", err)
		os.Exit(-1)
	}
}
