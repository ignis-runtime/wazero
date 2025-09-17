package sysfs_test

import (
	"io/fs"
	"testing/fstest"

	"github.com/ignis-runtime/wazero"
)

var moduleConfig wazero.ModuleConfig

// This example shows how to adapt a fs.FS to a sys.FS
func ExampleAdaptFS() {
	m := fstest.MapFS{
		"a/b.txt": &fstest.MapFile{Mode: 0o666},
		".":       &fstest.MapFile{Mode: 0o777 | fs.ModeDir},
	}
	root := &AdaptFS{FS: m}

	moduleConfig = wazero.NewModuleConfig().
		WithFSConfig(wazero.NewFSConfig().(FSConfig).WithSysFSMount(root, "/"))

	// Output:
}

// This example shows how to configure a sysfs.DirFS
func ExampleDirFS() {
	root := DirFS(".")

	moduleConfig = wazero.NewModuleConfig().
		WithFSConfig(wazero.NewFSConfig().(FSConfig).WithSysFSMount(root, "/"))

	// Output:
}

// This example shows how to configure a sysfs.ReadFS
func ExampleReadFS() {
	root := DirFS(".")
	readOnly := &ReadFS{FS: root}

	moduleConfig = wazero.NewModuleConfig().
		WithFSConfig(wazero.NewFSConfig().(FSConfig).WithSysFSMount(readOnly, "/"))

	// Output:
}
