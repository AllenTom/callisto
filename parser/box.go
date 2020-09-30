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
	Border        string                   `json:"border"`
	BorderTop     string                   `json:"borderTop"`
	BorderRight   string                   `json:"borderRight"`
	BorderBottom  string                   `json:"borderBottom"`
	BorderLeft    string                   `json:"borderLeft"`
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

	// parse border
	border, err := ParseBorder(docElm.Border)
	if err != nil {
		return nil, err
	}
	node.Border = border

	topWidth, topColor, err := ParseBorderWithPosition(docElm.BorderTop)
	if err != nil {
		return nil, err
	}
	node.Border.Top = callisto.Border{
		Width: topWidth,
		Color: topColor,
	}

	rightWidth, rightColor, err := ParseBorderWithPosition(docElm.BorderRight)
	if err != nil {
		return nil, err
	}
	node.Border.Right = callisto.Border{
		Width: rightWidth,
		Color: rightColor,
	}

	bottomWidth, bottomColor, err := ParseBorderWithPosition(docElm.BorderBottom)
	if err != nil {
		return nil, err
	}
	node.Border.Bottom = callisto.Border{
		Width: bottomWidth,
		Color: bottomColor,
	}

	leftWidth, leftColor, err := ParseBorderWithPosition(docElm.BorderLeft)
	if err != nil {
		return nil, err
	}
	node.Border.Left = callisto.Border{
		Width: leftWidth,
		Color: leftColor,
	}

	return node, nil
}
