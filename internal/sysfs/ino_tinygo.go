//go:build tinygo

package sysfs

import (
	"io/fs"

	experimentalsys "github.com/ASparkOfFire/wazero/experimental/sys"
	"github.com/ASparkOfFire/wazero/sys"
)

func inoFromFileInfo(_ string, info fs.FileInfo) (sys.Inode, experimentalsys.Errno) {
	return 0, experimentalsys.ENOTSUP
}
