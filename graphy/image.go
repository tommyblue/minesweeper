package graphy

import (
	"fmt"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

func (ui *sdlWrapper) warmCache() error {
	if ui.cache == nil {
		cache := make(map[string]imageStruct)
		ui.cache = &cache
	}

	for imageID, imagePath := range *ui.conf.ImagesToCache {
		if ui.conf.Debug {
			fmt.Printf("Caching image ID \"%s\"\n  File path: %s\n", imageID, imagePath)
		}
		image, err := img.Load(imagePath)
		if err != nil {
			return err
		}

		if ui.conf.Debug {
			fmt.Printf("  Image size: %dx%d\n", image.W, image.H)
		}
		surface, err := sdl.CreateRGBSurface(
			0,
			image.W,
			image.H,
			32,
			0xff000000,
			0x00ff0000,
			0x0000ff00,
			0x000000ff,
		)

		if err != nil {
			return err
		}

		srcRect := sdl.Rect{X: 0, Y: 0, W: image.W, H: image.H}

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

func (ui *sdlWrapper) drawImage(tile Tile) error {
	imageStruct := ui.getImageFromCache(tile.ImageID)
	x := imageStruct.rect.W * tile.PosX
	y := imageStruct.rect.H * tile.PosY

	dst := sdl.Rect{X: x, Y: y, W: imageStruct.rect.W, H: imageStruct.rect.H}
	err := ui.renderer.Copy(imageStruct.image, &imageStruct.rect, &dst)
	return err
}

func (ui *sdlWrapper) getImage(imageID string) imageStruct {
	return ui.getImageFromCache(imageID)
}

func (ui *sdlWrapper) getImageFromCache(imageID string) imageStruct {
	return (*ui.cache)[imageID]
}
