package game

import (
	"testing"
)

func TestCalculateNextState(t *testing.T) {
	tests := []struct {
		name     string
		initial  map[Cell]int
		expected map[Cell]int
	}{
		{
			name: "Head to Tail",
			initial: map[Cell]int{
				{0, 0}: ElectronHead,
			},
			expected: map[Cell]int{
				{0, 0}: ElectronTail,
			},
		},
		{
			name: "Tail to Conductor",
			initial: map[Cell]int{
				{0, 0}: ElectronTail,
			},
			expected: map[Cell]int{
				{0, 0}: Conductor,
			},
		},
		{
			name: "Conductor with 1 Head becomes Head",
			initial: map[Cell]int{
				{0, 0}: Conductor,
				{0, 1}: ElectronHead,
			},
			expected: map[Cell]int{
				{0, 0}: ElectronHead,
				{0, 1}: ElectronTail,
			},
		},
		{
			name: "Conductor with 2 Heads becomes Head",
			initial: map[Cell]int{
				{0, 0}: Conductor,
				{0, 1}: ElectronHead,
				{1, 0}: ElectronHead,
			},
			expected: map[Cell]int{
				{0, 0}: ElectronHead,
				{0, 1}: ElectronTail,
				{1, 0}: ElectronTail,
			},
		},
		{
			name: "Conductor with 3 Heads stays Conductor",
			initial: map[Cell]int{
				{0, 0}: Conductor,
				{0, 1}: ElectronHead,
				{1, 0}: ElectronHead,
				{1, 1}: ElectronHead,
			},
			expected: map[Cell]int{
				{0, 0}: Conductor,
				{0, 1}: ElectronTail,
				{1, 0}: ElectronTail,
				{1, 1}: ElectronTail,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := CalculateNextState(tc.initial)
			if len(result) != len(tc.expected) {
				t.Fatalf("expected len %d, got %d", len(tc.expected), len(result))
			}
			for k, v := range tc.expected {
				if result[k] != v {
					t.Errorf("cell %v expected %v, got %v", k, v, result[k])
				}
			}
		})
	}
}
