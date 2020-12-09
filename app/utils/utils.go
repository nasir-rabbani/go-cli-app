package utils

import "mycart/app/helpers/loghelper"

// HandelException - Logs error if any and condition based panic
func HandelException(err error, shouldPanic bool) {
	if err != nil {
		loghelper.LogError(err)
		if shouldPanic {
			panic(err)
		}
	}
}
