package testutil

import (
	"os"
	"path/filepath"
	"testing"
)

// WithCmdTestDir changes CWD to tmp/test/<testName> and returns a cleanup func.
// All relative artifacts (e.g., ./lc_base, ./lc_uploads) will be created inside it.
func WithCmdTestDir(t *testing.T) func() {
	t.Helper()
	oldwd, _ := os.Getwd()
	base := filepath.Join(oldwd, "tmp", "test", t.Name())
	if err := os.MkdirAll(base, 0o775); err != nil {
		t.Fatalf("failed to create %s: %v", base, err)
	}
	if err := os.Chdir(base); err != nil {
		t.Fatalf("failed to chdir to %s: %v", base, err)
	}
	return func() {
		_ = os.Chdir(oldwd)
		_ = os.RemoveAll(filepath.Join(oldwd, "tmp", "test", t.Name()))
	}
}
