package callisto

import "github.com/fogleman/gg"

type RenderContext struct {
	Context              *gg.Context
	ImageResourceLibrary *ImageResourceLibrary
	Parent               Element
	Sibling              Element
}
