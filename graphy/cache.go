package graphy

import (
	"fmt"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

// TODO: get from image
var imageSide int32 = 24

func (ui *sdlWrapper) warmCache() error {
	if ui.cache == nil {
		cache := make(map[string]imageStruct)
		ui.cache = &cache
	}

	for imageID, imagePath := range *ui.conf.ImagesToCache {
		fmt.Printf("Caching image ID \"%s\"\n", imageID)
		surface, err := sdl.CreateRGBSurface(
			0,
			imageSide,
			imageSide,
			32,
			0xff000000,
			0x00ff0000,
			0x0000ff00,
			0x000000ff,
		)

		if err != nil {
			return err
		}

		srcRect := sdl.Rect{X: 0, Y: 0, W: imageSide, H: imageSide}

		image, err := img.Load(imagePath)
		if err != nil {
			return err
		}

		err = image.Blit(&srcRect, surface, &srcRect)
		if err != nil {
			return err
		}

		imgTexture, err := ui.renderer.CreateTextureFromSurface(surface)
		if err != nil {
			return err
		}

		(*ui.cache)[imageID] = imageStruct{id: imageID, image: imgTexture, rect: srcRect}
	}
	return nil
}
