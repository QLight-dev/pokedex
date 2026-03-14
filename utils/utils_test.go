package utils

import "testing"

type tableTests struct {
	input float32
	want  float32
	name  string
}

func TestHectogramsToKg(t *testing.T) {
	tests := []tableTests{
		{name: "1 hectogram to kg", input: 1.0, want: 0.1},
		{name: "5 hectograms to kg", input: 5.0, want: 0.5},
		{name: "100 hectograms to kg", input: 100.0, want: 10.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := HectogramsToKg(tt.input)
			if got != tt.want {
				t.Errorf("input: %f, got %f, want: %f", tt.input, got, tt.want)
			}
		})
	}
}
