package utils

import "testing"

func TestInitializeLogger(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Intialize logger",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InitializeLogger()
		})
	}
}
