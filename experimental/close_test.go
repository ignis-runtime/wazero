package experimental_test

import (
	"context"
	"testing"

	"github.com/ignis-runtime/wazero/internal/expctxkeys"
	"github.com/ignis-runtime/wazero/internal/testing/require"
)

type arbitrary struct{}

// testCtx is an arbitrary, non-default context. Non-nil also prevents linter errors.
var testCtx = context.WithValue(context.Background(), arbitrary{}, "arbitrary")

func TestWithCloseNotifier(t *testing.T) {
	tests := []struct {
		name         string
		notification CloseNotifier
		expected     bool
	}{
		{
			name:     "returns input when notification nil",
			expected: false,
		},
		{
			name:         "decorates with notification",
			notification: CloseNotifyFunc(func(context.Context, uint32) {}),
			expected:     true,
		},
	}

	for _, tt := range tests {
		tc := tt
		t.Run(tc.name, func(t *testing.T) {
			if decorated := WithCloseNotifier(testCtx, tc.notification); tc.expected {
				require.NotNil(t, decorated.Value(expctxkeys.CloseNotifierKey{}))
			} else {
				require.Same(t, testCtx, decorated)
			}
		})
	}
}
