package v2

import (
	"context"
	"testing"

	"github.com/ignis-runtime/wazero"
	"github.com/ignis-runtime/wazero/api"
	"github.com/ignis-runtime/wazero/internal/integration_test/spectest"
	"github.com/ignis-runtime/wazero/internal/platform"
)

const enabledFeatures = api.CoreFeaturesV2

func TestCompiler(t *testing.T) {
	if !platform.CompilerSupported() {
		t.Skip()
	}
	spectest.Run(t, Testcases, context.Background(), wazero.NewRuntimeConfigCompiler().WithCoreFeatures(enabledFeatures))
}

func TestInterpreter(t *testing.T) {
	spectest.Run(t, Testcases, context.Background(), wazero.NewRuntimeConfigInterpreter().WithCoreFeatures(enabledFeatures))
}
