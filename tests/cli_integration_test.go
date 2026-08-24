package tests

import (
	"bytes"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestCLIIntegration(t *testing.T) {
	tempDir := t.TempDir()
	binPath := filepath.Join(tempDir, "echo-cli")

	buildCmd := exec.Command("go", "build", "-o", binPath, "../cmd/echo")
	out, err := buildCmd.CombinedOutput()
	if err != nil {
		t.Fatalf("failed to build echo-cli: %v, output: %s", err, string(out))
	}

	tests := []struct {
		name       string
		args       []string
		wantStdout string
	}{
		{
			name:       "row 1: echo-cli hello world",
			args:       []string{"hello", "world"},
			wantStdout: "hello\n",
		},
		{
			name:       "row 2: echo-cli hello",
			args:       []string{"hello"},
			wantStdout: "hello\n",
		},
		{
			name:       "row 3: echo-cli",
			args:       []string{},
			wantStdout: "\n",
		},
		{
			name:       "row 4: echo-cli \"hello world\" foo",
			args:       []string{"hello world", "foo"},
			wantStdout: "hello world\n",
		},
		{
			name:       "row 5: echo-cli \"\"",
			args:       []string{""},
			wantStdout: "\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := exec.Command(binPath, tt.args...)
			var stdout, stderr bytes.Buffer
			cmd.Stdout = &stdout
			cmd.Stderr = &stderr

			err := cmd.Run()
			if err != nil {
				t.Fatalf("unexpected execution error: %v", err)
			}

			if code := cmd.ProcessState.ExitCode(); code != 0 {
				t.Errorf("exit code = %d; want 0", code)
			}

			if gotStderr := stderr.String(); gotStderr != "" {
				t.Errorf("stderr = %q; want empty", gotStderr)
			}

			if gotStdout := stdout.String(); gotStdout != tt.wantStdout {
				t.Errorf("stdout = %q; want %q", gotStdout, tt.wantStdout)
			}
		})
	}
}
