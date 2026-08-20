package mousetrap

import (
	"runtime"
	"testing"
)

func TestStartedByExplorer(t *testing.T) {
	result := StartedByExplorer()
	if runtime.GOOS != "windows" && result != false {
		t.Errorf("expected StartedByExplorer() to return false on %s, got true", runtime.GOOS)
	}
}
