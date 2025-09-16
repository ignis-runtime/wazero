package wazevoapi

import (
	"testing"

	"github.com/ASparkOfFire/wazero/internal/testing/require"
)

func TestExitCode_withinByte(t *testing.T) {
	require.True(t, exitCodeMax < ExitCodeMask) //nolint
}
