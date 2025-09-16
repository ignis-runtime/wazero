package spectest

import (
	"context"
	"embed"
	"testing"

	"github.com/ASparkOfFire/wazero"
	"github.com/ASparkOfFire/wazero/api"
	"github.com/ASparkOfFire/wazero/experimental"
	"github.com/ASparkOfFire/wazero/internal/integration_test/spectest"
	"github.com/ASparkOfFire/wazero/internal/platform"
)

//go:embed testdata/*.wasm
//go:embed testdata/*.json
var testcases embed.FS

const enabledFeatures = api.CoreFeaturesV2 | experimental.CoreFeaturesThreads

func TestCompiler(t *testing.T) {
	if !platform.CompilerSupported() {
		t.Skip()
	}
	spectest.Run(t, testcases, context.Background(), wazero.NewRuntimeConfigCompiler().WithCoreFeatures(enabledFeatures))
}

func TestInterpreter(t *testing.T) {
	spectest.Run(t, testcases, context.Background(), wazero.NewRuntimeConfigInterpreter().WithCoreFeatures(enabledFeatures))
}
