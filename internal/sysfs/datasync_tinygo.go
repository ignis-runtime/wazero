//go:build tinygo

package sysfs

import (
	"os"

	"github.com/ASparkOfFire/wazero/experimental/sys"
)

func datasync(f *os.File) sys.Errno {
	return sys.ENOSYS
}
