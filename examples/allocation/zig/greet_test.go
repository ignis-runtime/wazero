package main

import (
	"testing"

	"github.com/ignis-runtime/wazero/internal/testing/maintester"
	"github.com/ignis-runtime/wazero/internal/testing/require"
)

// Test_main ensures the following will work:
//
//	go run greet.go wazero
func Test_main(t *testing.T) {
	stdout, _ := maintester.TestMain(t, main, "greet", "wazero")
	require.Equal(t, `wasm >> Hello, wazero!
go >> Hello, wazero!
`, stdout)
}
