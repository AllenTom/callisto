package parser

import (
	"github.com/allentom/callisto"
	"github.com/jinzhu/copier"
	"github.com/mitchellh/mapstructure"
)

type TextDocElement struct {
	Font  string  `json:"font"`
	Size  float64 `json:"size"`
	Text  string  `json:"text"`
	Color string  `json:"color"`
}

var TextHandler ElementHandler = func(context BuildContext) (callisto.Element, error) {
	boxElement, err := BoxHandler(context)
	if err != nil {
		return nil, err
	}
	box := boxElement.(*callisto.Box)
	var docElm TextDocElement
	err = mapstructure.Decode(&context.content, &docElm)
	if err != nil {
		return nil, err
	}

	node := callisto.TextBox{
		Box: *box,
	}
	err = copier.Copy(&node, &docElm)
	if err != nil {
		return nil, err
	}
	node.FontPath = ParseFont(docElm.Font)
	return &node, nil
}
