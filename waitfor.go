package chrono

import (
	"errors"
	"time"
)

type WaitForSettings struct {

	// Wait until this returns true, or MaxWaitTime expires
	Test func() bool
	MaxWaitTime time.Duration

	// How long to sleep between retries of the test func?
	SleepTime time.Duration

}

// A constructor for `WaitForSettings`
func MakeWaitForSettings(t func() bool) *WaitForSettings {
	return &WaitForSettings{
		Test:           t,
		SleepTime:   30 * time.Millisecond,
		MaxWaitTime: 10 * time.Second,
	}
}

// Convenience function that waits for the default settings
func WaitFor(w func() bool) error {
	return MakeWaitForSettings(w).Wait()
}

// A wrapper for WaitFor that panics upon error
func MustWaitFor(w func() bool) {
	err := WaitFor(w)
	if err != nil {
		panic(err)
	}
}

func (w *WaitForSettings) Wait() error {
	startTime := time.Now()
	for {
		time.Sleep(w.SleepTime)
		if w.Test() {
			break
		}
		if time.Since(startTime) > w.MaxWaitTime {
			return errors.New("Waited too long, giving up")
		}
	}
	return nil
}
