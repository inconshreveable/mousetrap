package mousetrap

import (
	"runtime"
	"testing"
)

func TestStartedByExplorer(t *testing.T) {
	result := StartedByExplorer()
	if runtime.GOOS != "windows" && result {
		t.Errorf("expected StartedByExplorer() to return false on %s, got true", runtime.GOOS)
	}
}

func TestIsExplorerOrFileManager(t *testing.T) {
	tests := []struct {
		name string
		exe  string
		want bool
	}{
		{"explorer lowercase", "explorer.exe", true},
		{"explorer uppercase", "EXPLORER.EXE", true},
		{"explorer mixed case", "Explorer.exe", true},
		{"explorer with windows path", `C:\Windows\explorer.exe`, true},
		{"explorer with forward slash path", "C:/Windows/explorer.exe", true},
		{"filepilot lowercase", "fpilot.exe", true},
		{"filepilot uppercase", "FPILOT.EXE", true},
		{"filepilot mixed case", "FPilot.exe", true},
		{"filepilot with path", `C:\Program Files\FilePilot\FPilot.exe`, true},
		{"filepilot with forward slash path", "C:/Program Files/FilePilot/FPilot.exe", true},
		{"cmd lowercase", "cmd.exe", false},
		{"cmd uppercase", "CMD.EXE", false},
		{"powershell", "powershell.exe", false},
		{"bash", "bash.exe", false},
		{"random binary", "mytool.exe", false},
		{"empty string", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isExplorerOrFileManager(tt.exe); got != tt.want {
				t.Errorf("isExplorerOrFileManager(%q) = %v, want %v", tt.exe, got, tt.want)
			}
		})
	}
}
