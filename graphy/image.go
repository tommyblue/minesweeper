package graphy

import (
	"errors"
	"fmt"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

func (ui *sdlWrapper) warmCache() error {
	if ui.cache == nil {
		cache := make(map[string]imageStruct)
		ui.cache = &cache
	}

	if ui.conf.BackgroundImage != "" {
		if ui.conf.Debug {
			fmt.Printf("Caching background\n  File path: %s\n", ui.conf.BackgroundImage)
		}
		imgTexture, srcRect, err := ui.getImageTexture(ui.conf.BackgroundImage)
		if err != nil {
			return err
		}
		(*ui.cache)["graphy-bg"] = imageStruct{id: "graphy-bg", image: imgTexture, rect: srcRect}
	}
	for imageID, imagePath := range *ui.conf.ImagesToCache {
		if ui.conf.Debug {
			fmt.Printf("Caching image ID \"%s\"\n  File path: %s\n", imageID, imagePath)
		}

		imgTexture, srcRect, err := ui.getImageTexture(imagePath)
		if err != nil {
			return err
		}
		(*ui.cache)[imageID] = imageStruct{id: imageID, image: imgTexture, rect: srcRect}
	}
	return nil
}

func (ui *sdlWrapper) getImageTexture(imagePath string) (*sdl.Texture, sdl.Rect, error) {
	image, err := img.Load(imagePath)
	if err != nil {
		return nil, sdl.Rect{}, err
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
		return nil, sdl.Rect{}, err
	}

	srcRect := sdl.Rect{X: 0, Y: 0, W: image.W, H: image.H}

	err = image.Blit(&srcRect, surface, &srcRect)
	if err != nil {
		return nil, sdl.Rect{}, err
	}

	imgTexture, err := ui.renderer.CreateTextureFromSurface(surface)
	if err != nil {
		return nil, sdl.Rect{}, err
	}

	return imgTexture, srcRect, nil
}

func (ui *sdlWrapper) drawBackground() error {
	return ui.drawImage(Tile{
		ImageID: "graphy-bg",
		PosX:    0,
		PosY:    0,
	}, nil)
}

func (ui *sdlWrapper) drawImage(tile Tile, offset *imageOffset) error {
	imageStruct, err := ui.getImageFromCache(tile.ImageID)
	if err != nil {
		return err
	}

	if offset == nil {
		offset = &imageOffset{x: 0, y: 0}
	}
	x := offset.x + imageStruct.rect.W*tile.PosX
	y := offset.y + imageStruct.rect.H*tile.PosY

	dst := sdl.Rect{X: x, Y: y, W: imageStruct.rect.W, H: imageStruct.rect.H}
	err = ui.renderer.Copy(imageStruct.image, &imageStruct.rect, &dst)
	return err
}

func (ui *sdlWrapper) getImage(imageID string) (imageStruct, error) {
	return ui.getImageFromCache(imageID)
}

func (ui *sdlWrapper) getImageFromCache(imageID string) (imageStruct, error) {
	img, ok := (*ui.cache)[imageID]
	if !ok {
		return imageStruct{}, errors.New("Can't find the image in the cache")
	}
	return img, nil
}
