package ui

import (
	"path"
	"runtime"
)

// return path to the file relative to the caller package
func getRelativePath(filepath string) string {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		panic("error")
	}
	return path.Join(path.Dir(filename), filepath)
}
