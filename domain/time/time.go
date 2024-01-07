package time

import (
	ti "time"
)

var mockedAt *ti.Time

// Stop stops the time at the set mock time, useful for deterministic `time.Now()` in tests.
func Stop(atTime ti.Time) {
	mockedAt = &atTime
}

// Restart clears a mocked time to reset behavior to core `time.Now()`
func Restart() {
	mockedAt = nil
}

// Now returns the mocked time if set. Else, `time.Now()`.
func Now() ti.Time {
	if mockedAt != nil {
		return *mockedAt
	}
	return ti.Now()
}

func IsStopped() bool {
	return mockedAt != nil
}
