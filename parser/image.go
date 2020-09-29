package parser

import (
	"github.com/allentom/callisto"
	"github.com/mitchellh/mapstructure"
)

type ImageDocElement struct {
	Src             string `json:"src"`
	OriginDimension bool   `json:"originDimension"`
	Fit             string `json:"fit"`
}

var ImageElementHandler ElementHandler = func(context BuildContext) (callisto.Element, error) {
	boxElement, err := BoxHandler(context)
	if err != nil {
		return nil, err
	}
	box := boxElement.(*callisto.Box)

	var docElm ImageDocElement
	err = mapstructure.Decode(&context.content, &docElm)
	if err != nil {
		return nil, err
	}

	node := callisto.ImageBox{
		Box:             *box,
		Src:             docElm.Src,
		OriginDimension: docElm.OriginDimension,
		Fit:             callisto.FitHeight,
	}

	imageFit, err := ParseImageFit(docElm.Fit)
	if err != nil {
		return nil, err
	}
	node.Fit = imageFit
	return &node, nil

}
