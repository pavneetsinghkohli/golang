package lib

import (
	"runtime"
)

func Run() string {
	return runtime.Version()
}
