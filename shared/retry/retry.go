package retry

import "errors"

var errMaxRetriesReached = errors.New("exceeded retry limit")

type Func func(attempt int) (retry bool, err error)

func Exec(maxRetries int, fn Func) error {
	var err error
	var cont bool
	attempt := 1
	for {
		cont, err = fn(attempt)
		if !cont || err == nil {
			break
		}
		attempt++
		if attempt > maxRetries {
			return errMaxRetriesReached
		}
	}
	return err
}
