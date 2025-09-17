package frontend

import (
	"github.com/ignis-runtime/wazero/internal/engine/wazevo/ssa"
	"github.com/ignis-runtime/wazero/internal/wasm"
)

func FunctionIndexToFuncRef(idx wasm.Index) ssa.FuncRef {
	return ssa.FuncRef(idx)
}
