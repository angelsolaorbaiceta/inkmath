package nums

import "testing"

func TestRangesOverlap(t *testing.T) {
	t.Run("Range always overlaps with itself", func(t *testing.T) {
		hasOverlap, start, end := RangesOverlap(1.0, 3.0, 1.0, 3.0)

		if !hasOverlap {
			t.Error("Expected ranges to overlap")
		}
		if start != 1.0 {
			t.Errorf("Wrong start. Want 1.0, got %f", start)
		}
		if end != 3.0 {
			t.Errorf("Wrong end. Want 3.0, got %f", end)
		}
	})

	t.Run("Range that contains other range overlaps", func(t *testing.T) {
		hasOverlap, start, end := RangesOverlap(1.0, 10.0, 3.0, 5.0)
		if !hasOverlap {
			t.Error("Expected ranges to overlap")
		}
		if start != 3.0 {
			t.Errorf("Wrong start. Want 3.0, got %f", start)
		}
		if end != 5.0 {
			t.Errorf("Wrong end. Want 5.0, got %f", end)
		}

		hasOverlap, start, end = RangesOverlap(3.0, 5.0, 1.0, 10.0)
		if !hasOverlap {
			t.Error("Expected ranges to overlap")
		}
		if start != 3.0 {
			t.Errorf("Wrong start. Want 3.0, got %f", start)
		}
		if end != 5.0 {
			t.Errorf("Wrong end. Want 5.0, got %f", end)
		}
	})

	t.Run("Overlapping ranges", func(t *testing.T) {
		hasOverlap, start, end := RangesOverlap(1.0, 10.0, 5.0, 15.0)
		if !hasOverlap {
			t.Error("Expected ranges to overlap")
		}
		if start != 5.0 {
			t.Errorf("Wrong start. Want 5.0, got %f", start)
		}
		if end != 10.0 {
			t.Errorf("Wrong end. Want 10.0, got %f", end)
		}

		hasOverlap, start, end = RangesOverlap(5.0, 15.0, 1.0, 10.0)
		if !hasOverlap {
			t.Error("Expected ranges to overlap")
		}
		if start != 5.0 {
			t.Errorf("Wrong start. Want 5.0, got %f", start)
		}
		if end != 10.0 {
			t.Errorf("Wrong end. Want 10.0, got %f", end)
		}
	})

	t.Run("Non overlapping ranges", func(t *testing.T) {
		if DoRangesOverlap(1.0, 10.0, 20.0, 25.0) {
			t.Error("Expected ranges to overlap")
		}
		if DoRangesOverlap(20.0, 25.0, 1.0, 10.0) {
			t.Error("Expected ranges to overlap")
		}
	})
}
