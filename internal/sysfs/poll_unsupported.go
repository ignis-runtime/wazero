//go:build !(linux || darwin || windows) || tinygo

package sysfs

import (
	"github.com/ignis-runtime/wazero/experimental/sys"
	"github.com/ignis-runtime/wazero/internal/fsapi"
)

// poll implements `Poll` as documented on fsapi.File via a file descriptor.
func poll(uintptr, fsapi.Pflag, int32) (bool, sys.Errno) {
	return false, sys.ENOSYS
}
