package mousetrap

import "strings"

// isExplorerOrFileManager reports whether the given executable name corresponds to
// Windows Explorer or a compatible file manager (such as FilePilot).
func isExplorerOrFileManager(exe string) bool {
	if i := strings.LastIndexAny(exe, `/\`); i >= 0 {
		exe = exe[i+1:]
	}
	return strings.EqualFold(exe, "explorer.exe") || strings.EqualFold(exe, "fpilot.exe")
}
