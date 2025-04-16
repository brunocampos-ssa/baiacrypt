package prompt

import (
	"fmt"
	"testing"
	"time"

	"github.com/creack/pty"
)

func TestReadPasswordWithConfirm_Success(t *testing.T) {
	ptmx, tty, err := pty.Open()
	if err != nil {
		t.Fatalf("failed to open pty: %v", err)
	}
	defer ptmx.Close()
	defer tty.Close()

	go func() {
		time.Sleep(100 * time.Millisecond)
		fmt.Fprint(ptmx, "test-password\n")
		time.Sleep(100 * time.Millisecond)
		fmt.Fprint(ptmx, "test-password\n")
	}()

	password, err := readPasswordWithConfirm(true, tty)
	if err != nil {
		t.Fatalf("readPasswordWithConfirm failed: %v", err)
	}

	expected := "test-password"
	if string(password) != expected {
		t.Errorf("expected %q, got %q", expected, password)
	}
}

func TestReadPasswordWithConfirm_Mismatch(t *testing.T) {
	ptmx, tty, err := pty.Open()
	if err != nil {
		t.Fatalf("failed to open pty: %v", err)
	}
	defer ptmx.Close()
	defer tty.Close()

	go func() {
		time.Sleep(100 * time.Millisecond)
		fmt.Fprint(ptmx, "one-password\n")
		time.Sleep(100 * time.Millisecond)
		fmt.Fprint(ptmx, "different-password\n")
	}()

	_, err = readPasswordWithConfirm(true, tty)
	if err == nil {
		t.Error("expected mismatch error but got none")
	}
}
