package mousetrap

import (
	"syscall"
	"unsafe"
)

var knownCallers = [2]string{"explorer.exe", "PowerToys.PowerLauncher.exe"}

func getProcessEntry(pid int) (*syscall.ProcessEntry32, error) {
	snapshot, err := syscall.CreateToolhelp32Snapshot(syscall.TH32CS_SNAPPROCESS, 0)
	if err != nil {
		return nil, err
	}
	defer syscall.CloseHandle(snapshot)
	var procEntry syscall.ProcessEntry32
	procEntry.Size = uint32(unsafe.Sizeof(procEntry))
	if err = syscall.Process32First(snapshot, &procEntry); err != nil {
		return nil, err
	}
	for {
		if procEntry.ProcessID == uint32(pid) {
			return &procEntry, nil
		}
		err = syscall.Process32Next(snapshot, &procEntry)
		if err != nil {
			return nil, err
		}
	}
}

// StartedByExplorer returns true if the program was invoked by the user double-clicking
// on the executable from explorer.exe or by using "PowerToys Run"
//
// It is conservative and returns false if any of the internal calls fail.
// It does not guarantee that the program was run from a terminal. It only can tell you
// whether it was launched from explorer.exe or "PowerToys Run"
func StartedByExplorer() bool {
	pe, err := getProcessEntry(syscall.Getppid())
	if err != nil {
		return false
	}
	caller := syscall.UTF16ToString(pe.ExeFile[:])
	for _, exe := range &knownCallers {
		if exe == caller {
			return true
		}
	}
	return false
}
