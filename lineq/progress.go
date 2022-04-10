package lineq

import (
	"math"
)

type IterativeSolverProgress struct {
	ProgressPercentage int
	Error              float64
	IterCount          int
}

// computeProgressPercentage returns the percentage of the progress or the current error.
// An error equal to, or smaller than, the maxError will return 100%.
// Every order of magnitude larger than the maxError will return 10% of the maxError, thus
// the progress is logarithmic.
func computeProgressPercentage(maxError, currentError float64) int {
	diff := math.Log10(currentError) - math.Log10(maxError)
	capDiff := math.Max(math.Min(diff, 10), 0)

	return int(10 * (10 - capDiff))
}

type computeProgressRequest struct {
	currentErrorFn func() float64
	iterCount      int
	maxError       float64
}

// computeProgress receives tasks to compute the progress percentage and the current error.
// The progress percentage is computed based on the maxError and the current error.
// The result is sent to the output channel, and when the input channel is exhausted,
// it closes the output channel.
func computeProgress(in <-chan computeProgressRequest, out chan<- IterativeSolverProgress) {
	lastProgressPercentage := -1

	for req := range in {
		var (
			currentError       = req.currentErrorFn()
			progressPercentage = computeProgressPercentage(req.maxError, currentError)
		)

		if progressPercentage != lastProgressPercentage {
			lastProgressPercentage = progressPercentage

			out <- IterativeSolverProgress{
				ProgressPercentage: progressPercentage,
				Error:              currentError,
				IterCount:          req.iterCount,
			}
		}
	}

	close(out)
}
