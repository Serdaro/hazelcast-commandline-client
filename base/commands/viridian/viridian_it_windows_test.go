//go:build windows && (std || viridian)

package viridian_test

import (
	"testing"
)

func streamLogs_NonInteractiveTest(t *testing.T) {
	t.Skipf("This test doesn't run on Windows")
}
