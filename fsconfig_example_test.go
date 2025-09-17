package wazero_test

import (
	"embed"
	"io/fs"
	"log"
)

//go:embed testdata/index.html
var testdataIndex embed.FS

var moduleConfig ModuleConfig

// This example shows how to configure an embed.FS.
func Example_fsConfig() {
	// Strip the embedded path testdata/
	rooted, err := fs.Sub(testdataIndex, "testdata")
	if err != nil {
		log.Panicln(err)
	}

	moduleConfig = NewModuleConfig().
		// Make "index.html" accessible to the guest as "/index.html".
		WithFSConfig(NewFSConfig().WithFSMount(rooted, "/"))

	// Output:
}
