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

/* Center the position of the image in a loop of images to be drawn.
The idea is that the game should draw a list of images, like buttons, one below the other,
with a given margin between them. This currently supports images drawn in a vertical row.
Params:
	- imgWidth: Width of the image
	- imgHeight: Height of the image
	- imgIndex: Index of the image in the loop of images to draw
	- imgsNumber: Total number of images to draw
*/

func centerImgPosition(imgWidth, imgHeight, imgIndex, imgsNumber int32) (int32, int32) {
	var margin int32 = 4
	var windowW int32 = 800
	var windowH int32 = 600
	x := (windowW / 2) - (imgWidth / 2)
	y := 4*int32(imgIndex) + (windowH / 2) - (imgHeight / 2) - ((int32(imgsNumber)*imgHeight + int32(margin*(imgsNumber-1))) / 2)
	return x, y
}
