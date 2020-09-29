package parser

import (
	"github.com/allentom/callisto"
	"github.com/jinzhu/copier"
	"github.com/mitchellh/mapstructure"
)

type ParagraphDocElement struct {
	Text  string  `json:"text"`
	Font  string  `json:"font"`
	Size  float64 `json:"size"`
	Color string  `json:"color"`
	Wrap  bool    `json:"wrap"`
}

var ParagraphHandler ElementHandler = func(context BuildContext) (callisto.Element, error) {
	boxElement, err := BoxHandler(context)
	if err != nil {
		return nil, err
	}
	box := boxElement.(*callisto.Box)
	var docElm ParagraphDocElement
	err = mapstructure.Decode(&context.content, &docElm)
	if err != nil {
		return nil, err
	}

	node := callisto.ParagraphBox{
		Box: *box,
	}
	err = copier.Copy(&node, &docElm)
	if err != nil {
		return nil, err
	}
	node.FontPath = ParseFont(docElm.Font)
	return &node, nil
}
