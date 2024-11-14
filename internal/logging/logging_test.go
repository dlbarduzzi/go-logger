package logging

import "testing"

func TestNewLogget(t *testing.T) {
	testCases := []struct {
		mode  string
		level string
	}{
		{mode: "dev", level: "debug"},
		{mode: "prod", level: "debug"},
	}

	for _, tc := range testCases {
		t.Run(tc.mode, func(t *testing.T) {
			t.Parallel()
			logger := NewLogger(tc.mode, tc.level)
			if logger == nil {
				t.Fatal("expected logger not to be nil")
			}
		})
	}
}
