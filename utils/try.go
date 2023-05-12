package utils

import (
	"errors"
	"time"
)

// MaxRetries 最大允许重试次数
var MaxRetries = 3

var errMaxRetriesReached = errors.New("exceeded retry limit")

// Func represents functions that can be retried.
type Func func(attempt int) (retry bool, err error)

// Do keeps trying the function until the second argument
// returns false, or no error is returned.
func Do(fn Func) (err error) {
	var cont bool
	attempt := 1
	for {
		if attempt > 1 {
			time.Sleep(2 * time.Second)
		}
		cont, err = fn(attempt)
		if !cont || err == nil {
			break
		}
		attempt++
		if attempt > MaxRetries {
			return errMaxRetriesReached
		}
	}

	return
}
