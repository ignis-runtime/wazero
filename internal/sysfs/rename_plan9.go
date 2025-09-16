package sysfs

import (
	"os"

	"github.com/ASparkOfFire/wazero/experimental/sys"
)

func rename(from, to string) sys.Errno {
	if from == to {
		return 0
	}
	return sys.UnwrapOSError(os.Rename(from, to))
}
