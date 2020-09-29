package callisto

import (
	"image"
	"os"
)

type ImageElement struct {
	ImagePath string
	Content   *image.Image
	Config    *image.Config
}

func NewImageElement(imagePath string) (*ImageElement, error) {
	var err error
	elm := &ImageElement{ImagePath: imagePath}
	elm.Content, err = ReadImageFromFile(imagePath)
	if err != nil {
		return nil, err
	}
	elm.Config, err = ReadImageDimension(imagePath)
	if err != nil {
		return nil, err
	}
	return elm, nil
}

func ReadImageFromFile(imageFilePath string) (*image.Image, error) {
	f, err := os.Open(imageFilePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}
	return &img, nil
}

func ReadImageDimension(imagePath string) (*image.Config, error) {
	file, err := os.Open(imagePath)
	if err != nil {
		return nil, err
	}

	config, _, err := image.DecodeConfig(file)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
