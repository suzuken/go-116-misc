package testfailedonexit0

import (
	"os"
	"testing"
)

// This will cause below:
//
// === RUN   TestFunction
// --- FAIL: TestFunction (0.00s)
// panic: unexpected call to os.Exit(0) during test [recovered]
//         panic: unexpected call to os.Exit(0) during test
func TestFunction(t *testing.T) {
	os.Exit(0)
}
