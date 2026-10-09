package mousetrap

import "strings"

// KnownFileManagers lists process executable names recognized as Windows file managers.
// Applications can inspect or append additional file manager names to this slice.
var KnownFileManagers = []string{
	"explorer.exe",
	"fpilot.exe",
	"filepilot.exe",
	"totalcmd.exe",
	"totalcmd64.exe",
	"dopus.exe",
	"xyplorer.exe",
	"freecommander.exe",
	"onecommander.exe",
}

// AddFileManager registers an additional file manager executable name to KnownFileManagers.
func AddFileManager(name string) {
	name = strings.TrimSpace(name)
	if name != "" {
		KnownFileManagers = append(KnownFileManagers, name)
	}
}

// isExplorerOrFileManager reports whether the given executable name corresponds to
// Windows Explorer or an alternative file manager (such as FilePilot).
func isExplorerOrFileManager(exe string) bool {
	exe = strings.TrimSpace(exe)
	if i := strings.LastIndexAny(exe, `/\`); i >= 0 {
		exe = exe[i+1:]
	}
	exe = strings.TrimSpace(exe)
	if exe == "" {
		return false
	}
	for _, manager := range KnownFileManagers {
		if strings.EqualFold(exe, manager) {
			return true
		}
	}
	return false
}
