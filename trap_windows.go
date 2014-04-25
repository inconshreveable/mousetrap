// +build windows
package mousetrap

import (
    "os"
    "fmt"
    "syscall"
    "strings"
    "unsafe"
)

const (
    th32cs_snapprocess uintptr = 0x2
)

type processEntry32 struct {
    dwSize uint32
    cntUsage uint32
    th32ProcessID uint32
    th32DefaultHeapID int
    th32ModuleID uint32
    cntThreads uint32
    th32ParentProcessID uint32
    pcPriClassBase int32
    dwFlags uint32
    szExeFile [0x00000104]byte
}

func getProcessEntry(pid int) (pe *processEntry32, err error) {
    kernel := syscall.MustLoadDLL("kernel32.dll")
    CreateToolhelp32Snapshot := kernel.MustFindProc("CreateToolhelp32Snapshot")
    Process32First := kernel.MustFindProc("Process32First")
    Process32Next := kernel.MustFindProc("Process32Next")
    CloseHandle := kernel.MustFindProc("CloseHandle")

    snapshot, _, e1 := CreateToolhelp32Snapshot.Call(th32cs_snapprocess, uintptr(0));
    if (snapshot == uintptr(syscall.InvalidHandle)) {
        err = fmt.Errorf("CreateToolhelp32Snapshot: %v", e1)
        return
    }
    defer CloseHandle.Call(snapshot)

    var processEntry processEntry32
    processEntry.dwSize = uint32(unsafe.Sizeof(processEntry))
    ok, _ , e1 := Process32First.Call(snapshot, uintptr(unsafe.Pointer(&processEntry)))
    if ok == 0 {
        err = fmt.Errorf("Process32First: %v", e1)
        return
    }

    for {
        if processEntry.th32ProcessID == uint32(pid) {
            pe = &processEntry
            return
        }

        ok, _, e1 = Process32Next.Call(snapshot, uintptr(unsafe.Pointer(&processEntry)))
        if ok == 0 {
            err = fmt.Errorf("Process32Next: %v", e1)
            return
        }
    }
}

func getppid() (pid int, err error) {
    pe, err := getProcessEntry(os.Getpid())
    if err != nil {
        return
    }

    pid = int(pe.th32ParentProcessID)
    return
}

func InvokedFromCommandLine() (bool, error) {
    ppid, err := getppid()
    if err != nil {
        return true, err
    }

    pe, err := getProcessEntry(ppid)
    if err != nil {
        return true, err
    }

    var path string
    for i, b := range pe.szExeFile[:] {
        if b == 0 {
            path = string(pe.szExeFile[:i])
            break
        }
    }

    isExplorer := strings.HasSuffix(path, "explorer.exe")
    return !isExplorer, nil
}
