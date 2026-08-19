//go:build windows
// +build windows

package mousetrap

import "testing"

func TestStartedByExplorerWindows(t *testing.T) {
	_ = StartedByExplorer()
}
