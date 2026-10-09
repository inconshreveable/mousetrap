//go:build windows
// +build windows

package mousetrap

import (
	"os"
	"testing"
)

func TestGetProcessEntry(t *testing.T) {
	pe, err := getProcessEntry(os.Getpid())
	if err != nil {
		t.Fatalf("getProcessEntry(os.Getpid()) returned unexpected error: %v", err)
	}
	if pe == nil {
		t.Fatal("expected non-nil ProcessEntry32")
	}
	if int(pe.ProcessID) != os.Getpid() {
		t.Errorf("expected ProcessID %d, got %d", os.Getpid(), pe.ProcessID)
	}
}

func TestGetProcessEntryInvalidPID(t *testing.T) {
	_, err := getProcessEntry(-1)
	if err == nil {
		t.Error("expected error for PID -1, got nil")
	}
}
