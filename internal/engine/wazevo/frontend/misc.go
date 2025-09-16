package frontend

import (
	"github.com/ASparkOfFire/wazero/internal/engine/wazevo/ssa"
	"github.com/ASparkOfFire/wazero/internal/wasm"
)

func FunctionIndexToFuncRef(idx wasm.Index) ssa.FuncRef {
	return ssa.FuncRef(idx)
}
