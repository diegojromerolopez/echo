package echoer

import (
	"testing"
)

func TestEcho(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		expected string
	}{
		{
			name:     "multiple args",
			args:     []string{"hello", "world"},
			expected: "hello",
		},
		{
			name:     "single arg",
			args:     []string{"only-one"},
			expected: "only-one",
		},
		{
			name:     "empty slice",
			args:     []string{},
			expected: "",
		},
		{
			name:     "nil slice",
			args:     nil,
			expected: "",
		},
		{
			name:     "empty string first arg",
			args:     []string{""},
			expected: "",
		},
		{
			name:     "whitespace preservation",
			args:     []string{"hello world", "foo"},
			expected: "hello world",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Echo(tt.args)
			if got != tt.expected {
				t.Errorf("Echo(%v) = %q; want %q", tt.args, got, tt.expected)
			}
		})
	}
}
