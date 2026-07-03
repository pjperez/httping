package main

import (
	"math"
	"testing"
)

// TestCalculateMinMaxGuardsEmpty locks in the empty-slice guard added during
// the code review so calculateMinMax can never panic on an empty input.
func TestCalculateMinMaxGuardsEmpty(t *testing.T) {
	min, max := calculateMinMax(nil)
	if min != 0 || max != 0 {
		t.Fatalf("expected (0,0) for empty input, got (%v,%v)", min, max)
	}
}

func TestCalculateMinMaxNormal(t *testing.T) {
	min, max := calculateMinMax([]float64{3.0, 1.0, 2.0})
	if min != 1.0 || max != 3.0 {
		t.Fatalf("expected (1,3), got (%v,%v)", min, max)
	}
}

// TestPercentileDurationNaN verifies that a failed/NaN percentile result is
// rendered as "N/A" instead of being converted into a garbage time.Duration
// (the bug that previously produced values like "-2562047h47m16.854775808s").
func TestPercentileDurationNaN(t *testing.T) {
	got := percentileDuration(0, errFailed)
	if got != "N/A" {
		t.Fatalf("expected N/A on error, got %v", got)
	}
	// NaN check
	got = percentileDuration(math.NaN(), nil)
	if got != "N/A" {
		t.Fatalf("expected N/A for NaN, got %v", got)
	}
}

var errFailed = errFailedVal{}

type errFailedVal struct{}

func (errFailedVal) Error() string { return "failed" }
