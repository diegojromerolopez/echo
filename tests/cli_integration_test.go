package tests

import (
	"os/exec"
	"path/filepath"
	"testing"
)

func TestCLI(t *testing.T) {
	tempDir := t.TempDir()
	binPath := filepath.Join(tempDir, "echo-cli")

	buildCmd := exec.Command("go", "build", "-o", binPath, "../cmd/echo")
	out, err := buildCmd.CombinedOutput()
	if err != nil {
		t.Fatalf("failed to build echo-cli: %v, output: %s", err, string(out))
	}

	t.Run("with args", func(t *testing.T) {
		cmd := exec.Command(binPath, "hello", "world")
		stdout, err := cmd.Output()
		if err != nil {
			t.Fatalf("execution failed: %v", err)
		}
		if got := string(stdout); got != "hello\n" {
			t.Errorf("got %q, want %q", got, "hello\n")
		}
	})

	t.Run("with whitespace arg", func(t *testing.T) {
		cmd := exec.Command(binPath, "hello world", "foo")
		stdout, err := cmd.Output()
		if err != nil {
			t.Fatalf("execution failed: %v", err)
		}
		if got := string(stdout); got != "hello world\n" {
			t.Errorf("got %q, want %q", got, "hello world\n")
		}
	})

	t.Run("without args", func(t *testing.T) {
		cmd := exec.Command(binPath)
		stdout, err := cmd.Output()
		if err != nil {
			t.Fatalf("execution failed: %v", err)
		}
		if got := string(stdout); got != "\n" {
			t.Errorf("got %q, want %q", got, "\n")
		}
	})
}
