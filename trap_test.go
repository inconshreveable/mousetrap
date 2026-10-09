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
		{
			name: "explorer lowercase",
			exe:  "explorer.exe",
			want: true,
		},
		{
			name: "explorer uppercase",
			exe:  "EXPLORER.EXE",
			want: true,
		},
		{
			name: "explorer mixed case",
			exe:  "Explorer.exe",
			want: true,
		},
		{
			name: "explorer windows backslash path",
			exe:  `C:\Windows\explorer.exe`,
			want: true,
		},
		{
			name: "explorer forward slash path",
			exe:  "C:/Windows/explorer.exe",
			want: true,
		},
		{
			name: "explorer device path",
			exe:  `\Device\HarddiskVolume3\Windows\explorer.exe`,
			want: true,
		},
		{
			name: "filepilot short name lowercase",
			exe:  "fpilot.exe",
			want: true,
		},
		{
			name: "filepilot short name mixed case",
			exe:  "FPilot.exe",
			want: true,
		},
		{
			name: "filepilot short name uppercase",
			exe:  "FPILOT.EXE",
			want: true,
		},
		{
			name: "filepilot full name lowercase",
			exe:  "filepilot.exe",
			want: true,
		},
		{
			name: "filepilot full name mixed case",
			exe:  "FilePilot.exe",
			want: true,
		},
		{
			name: "filepilot with windows backslash path",
			exe:  `C:\Program Files\FilePilot\FPilot.exe`,
			want: true,
		},
		{
			name: "filepilot with forward slash path",
			exe:  "C:/Program Files/FilePilot/FPilot.exe",
			want: true,
		},
		{
			name: "totalcmd 32-bit",
			exe:  "totalcmd.exe",
			want: true,
		},
		{
			name: "totalcmd 64-bit uppercase",
			exe:  "TOTALCMD64.EXE",
			want: true,
		},
		{
			name: "directory opus",
			exe:  "dopus.exe",
			want: true,
		},
		{
			name: "xyplorer",
			exe:  "XYPlorer.exe",
			want: true,
		},
		{
			name: "freecommander",
			exe:  "FreeCommander.exe",
			want: true,
		},
		{
			name: "onecommander",
			exe:  "OneCommander.exe",
			want: true,
		},
		{
			name: "cmd lowercase",
			exe:  "cmd.exe",
			want: false,
		},
		{
			name: "cmd uppercase",
			exe:  "CMD.EXE",
			want: false,
		},
		{
			name: "powershell",
			exe:  "powershell.exe",
			want: false,
		},
		{
			name: "pwsh",
			exe:  "pwsh.exe",
			want: false,
		},
		{
			name: "bash",
			exe:  "bash.exe",
			want: false,
		},
		{
			name: "windows terminal",
			exe:  "wt.exe",
			want: false,
		},
		{
			name: "console host",
			exe:  "conhost.exe",
			want: false,
		},
		{
			name: "arbitrary executable",
			exe:  "mytool.exe",
			want: false,
		},
		{
			name: "prefixed name is not a match",
			exe:  "not_explorer.exe",
			want: false,
		},
		{
			name: "suffixed extension is not a match",
			exe:  "explorer.exe.bak",
			want: false,
		},
		{
			name: "suffixed filepilot is not a match",
			exe:  "fpilot.exe.old",
			want: false,
		},
		{
			name: "empty string",
			exe:  "",
			want: false,
		},
		{
			name: "whitespace only",
			exe:  "   ",
			want: false,
		},
		{
			name: "whitespace padded explorer",
			exe:  "  explorer.exe  ",
			want: true,
		},
		{
			name: "whitespace padded filepilot",
			exe:  "  FPilot.exe\n",
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isExplorerOrFileManager(tt.exe); got != tt.want {
				t.Errorf("isExplorerOrFileManager(%q) = %v, want %v", tt.exe, got, tt.want)
			}
		})
	}
}

func TestAddFileManager(t *testing.T) {
	saved := make([]string, len(KnownFileManagers))
	copy(saved, KnownFileManagers)
	t.Cleanup(func() {
		KnownFileManagers = saved
	})

	customManager := "custompilot.exe"
	if isExplorerOrFileManager(customManager) {
		t.Fatalf("expected %q not to be recognized initially", customManager)
	}

	AddFileManager(customManager)

	if !isExplorerOrFileManager(customManager) {
		t.Errorf("expected %q to be recognized after AddFileManager", customManager)
	}

	// Verify case-insensitive matching for dynamically added managers
	if !isExplorerOrFileManager("CUSTOMPILOT.EXE") {
		t.Errorf("expected uppercase CUSTOMPILOT.EXE to be recognized")
	}

	// Verify path-prefixed matching for dynamically added managers
	if !isExplorerOrFileManager(`D:\Tools\custompilot.exe`) {
		t.Errorf("expected path-prefixed custompilot.exe to be recognized")
	}

	// Blank string should be ignored
	initialLen := len(KnownFileManagers)
	AddFileManager("   ")
	if len(KnownFileManagers) != initialLen {
		t.Errorf("expected length %d after adding whitespace, got %d", initialLen, len(KnownFileManagers))
	}
}
