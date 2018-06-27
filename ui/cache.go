package ui

import (
	"fmt"
)

// returns a map with key of the image and its path
func getImagesToCache() *map[string]string {
	ret := make(map[string]string)
	for _, v := range tileImages {
		ret[string(v)] = getAbsolutePath(fmt.Sprintf("../assets/images/%s.png", v))
	}
	return &ret
}
