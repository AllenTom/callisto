package parser

import (
	"github.com/allentom/callisto"
	"github.com/mitchellh/mapstructure"
)

type BoxDocElement struct {
	Width         string                   `json:"width"`
	Height        string                   `json:"height"`
	LayoutManager string                   `json:"layoutManager"`
	ChildAlign    string                   `json:"childAlign"`
	Background    string                   `json:"background"`
	Children      []map[string]interface{} `json:"children"`
	Padding       string                   `json:"padding"`
	BoxSizing     string                   `json:"boxSizing"`
}

var BoxHandler ElementHandler = func(context BuildContext) (callisto.Element, error) {
	var docElm BoxDocElement
	err := mapstructure.Decode(context.content, &docElm)
	if err != nil {
		return nil, err
	}
	node := &callisto.Box{
		// parse background
		Color: docElm.Background,
		//parse layout manager
		LayoutManager: GetLayoutManagerWithString(docElm.LayoutManager),
		// parse child align
		ChildrenAlign: GetChildAlignWithString(docElm.ChildAlign),
		ElementPosition: callisto.ElementPosition{
			UseParentDelta: true,
		},
	}

	//parse dimension
	dimension := callisto.ElementDimension{}

	//parse width
	width, widthUnit, err := ParseValueAndUnit(docElm.Width)
	switch widthUnit {
	case ParentScaleRelative:
		dimension.UserParentRelative = true
		dimension.ParentRelativeScaleWidth = width / 100
	case Pixel:
		dimension.Width = width
	case ParentDeltaRelative:
		dimension.UserParentRelative = true
		dimension.ParentRelativeDeltaWidth = width
	}

	//parse height
	height, heightUnit, err := ParseValueAndUnit(docElm.Height)
	switch heightUnit {
	case ParentScaleRelative:
		dimension.UserParentRelative = true
		dimension.ParentRelativeScaleHeight = height / 100
	case Pixel:
		dimension.Height = height
	case ParentDeltaRelative:
		dimension.UserParentRelative = true
		dimension.ParentRelativeDeltaWidth = height
	}
	node.ElementDimension = dimension

	// parse padding
	node.Padding = GetPaddingFromRaw(docElm.Padding)

	//parse boxSizing
	node.BoxSizing = ParseBoxSizing(docElm.BoxSizing)
	return node, nil
}
