package lineq

import (
	"fmt"
	"testing"
)

func TestComputeProgress(t *testing.T) {
	maxError := 0.0001

	cases := []struct {
		currentError   float64
		wantPercentage int
	}{
		{0.00001, 100},
		{0.0001, 100},
		{0.001, 90},
		{0.01, 80},
		{0.1, 70},
		{1.0, 60},
		{10.0, 50},
		{100.0, 40},
		{1000.0, 30},
		{10000.0, 20},
		{100000.0, 10},
		{1000000.0, 0},
		{10000000.0, 0},
	}

	for _, testCase := range cases {
		description := fmt.Sprintf(
			"when the error is %f, the percentage should be %d",
			testCase.currentError, testCase.wantPercentage,
		)

		t.Run(description, func(t *testing.T) {
			got := computeProgressPercentage(maxError, testCase.currentError)
			if got != testCase.wantPercentage {
				t.Errorf("got %d, want %d", got, testCase.wantPercentage)
			}
		})
	}
}
