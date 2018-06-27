package ui

import (
	"path"
	"runtime"
)

func getAbsolutePath(filepath string) string {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		panic("error")
	}
	return path.Join(path.Dir(filename), filepath)
}
