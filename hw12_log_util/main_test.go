package main

import (
	"testing"
)

// TestLogLevelStatistic tests the logLevelStatistic function.
func TestLogLevelStatistic(t *testing.T) {
	tests := []struct {
		name     string
		data     []byte
		logLevel string
		want     int
	}{
		{
			name:     "2",
			data:     []byte("info: message 1\ndebug: message 2\ninfo: message 3"),
			logLevel: "info",
			want:     2,
		},
		{
			name:     "0",
			data:     []byte("info: message 1\ndebug: message 2\ninfo: message 3"),
			logLevel: "error",
			want:     0,
		},
		{
			name:     "3",
			data:     []byte("debug: message 1\ndebug: message 2\ndebug: message 3"),
			logLevel: "debug",
			want:     3,
		},
		{
			name:     "Empty Data",
			data:     []byte(""),
			logLevel: "info",
			want:     0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := logLevelStatistic(tt.data, tt.logLevel); got != tt.want {
				t.Errorf("logLevelStatistic() = %v, want %v", got, tt.want)
			}
		})
	}
}
