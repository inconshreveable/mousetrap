//go:build !windows
// +build !windows

package mousetrap

import "testing"

func TestStartedByExplorerNonWindows(t *testing.T) {
	if StartedByExplorer() {
		t.Errorf("StartedByExplorer() on non-Windows should return false, got true")
	}
}
