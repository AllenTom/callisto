package callisto

type ImageResourceLibrary struct {
	Images []*ImageElement
}

func (l *ImageResourceLibrary) LoadFromResource(imageFilePath string) error {
	element, err := NewImageElement(imageFilePath)
	if err != nil {
		return err
	}
	l.Images = append(l.Images, element)
	return nil
}

func (l *ImageResourceLibrary) GetImageWithPath(imagePath string) (target *ImageElement) {
	for _, element := range l.Images {
		if element.ImagePath == imagePath {
			target = element
			return
		}
	}
	return
}
