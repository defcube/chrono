package chrono

import (
	"errors"
	"time"
)

// WaitForSettings gives low-level control. Generally, one would call
// `MakeWaitForSettings` to initialize this to default settings
type WaitForSettings struct {

	// Wait until this returns true, or MaxWaitTime expires
	Test func() bool
	MaxWaitTime time.Duration

	// How long to sleep between retries of the test func?
	SleepTime time.Duration

}

// MakeWaitForSettings returns a `WaitForSettings` struct with fields
// initialized to defaults
func MakeWaitForSettings(t func() bool) *WaitForSettings {
	return &WaitForSettings{
		Test:           t,
		SleepTime:   30 * time.Millisecond,
		MaxWaitTime: 10 * time.Second,
	}
}

// WaitFor blocks until the condition returns True
func WaitFor(w func() bool) error {
	return MakeWaitForSettings(w).Wait()
}

// MustWaitFor is a wrapper for WaitFor that panics upon error
func MustWaitFor(w func() bool) {
	err := WaitFor(w)
	if err != nil {
		panic(err)
	}
}

// Wait blocks until the condition returns True
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
