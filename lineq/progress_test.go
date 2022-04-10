package lineq

import (
	"fmt"
	"testing"
)

func TestComputeProgressPercentage(t *testing.T) {
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

func TestComputeProgress(t *testing.T) {
	var (
		maxError = 0.0001
		inChan   = make(chan computeProgressRequest)
		outChan  = make(chan IterativeSolverProgress)
	)

	go computeProgress(inChan, outChan)

	t.Run("when progress is notified for the first time", func(t *testing.T) {
		inChan <- computeProgressRequest{
			currentErrorFn: func() float64 { return 100.0 },
			iterCount:      0,
			maxError:       maxError,
		}

		if got := <-outChan; got.ProgressPercentage != 40 {
			t.Errorf("want %d%%, got %d%%", 40, got.ProgressPercentage)
		}
	})

	t.Run("when the same error is notified, it doesn't emit progress", func(t *testing.T) {
		inChan <- computeProgressRequest{
			currentErrorFn: func() float64 { return 100.0 },
			iterCount:      1,
			maxError:       maxError,
		}

		select {
		case progress := <-outChan:
			t.Errorf("got progress %v (expected nothing)", progress)
		default:
			break
		}
	})

	t.Run("when a new error is notified, it emits progress", func(t *testing.T) {
		inChan <- computeProgressRequest{
			currentErrorFn: func() float64 { return 10.0 },
			iterCount:      2,
			maxError:       maxError,
		}

		if got := <-outChan; got.ProgressPercentage != 50 {
			t.Errorf("want %d%%, got %d%%", 50, got.ProgressPercentage)
		}
	})

	close(inChan)
}
