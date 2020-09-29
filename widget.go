package callisto

import (
	"fmt"
)

type RenderTask struct {
	Element     Element
	GetChildren []Element
}
type PreloadElement interface {
	Preload(ctx *RenderContext)
}

type Element interface {
	Position
	Dimension
	Render(context *RenderContext)
	GetChildren() []Element
	GetAnchor(anchor Anchor) (X, Y float64)
	AlignWithAnchor(anchor Anchor, X float64, Y float64)
	ApplyChildPosition()
	AddToChildren(element Element)
}

type ElementPadding struct {
	Left   float64
	Right  float64
	Top    float64
	Bottom float64
}

type BoxSizing int

const (
	BorderBox BoxSizing = iota + 1
	ContentBox
)

type Box struct {
	ElementPosition
	ElementDimension
	Children      []Element
	ChildrenAlign ChildAlign
	Color         string
	LayoutManager LayoutManager
	DrawArea      DrawArea
	Padding       ElementPadding
	BoxSizing     BoxSizing
}

func (b *Box) TransformChildPosition() {
	b.DrawArea.X = b.X + b.Padding.Left
	b.DrawArea.Y = b.Y + b.Padding.Top
	if b.LayoutManager != nil {
		b.LayoutManager.Place(b)
	}
	if b.ChildrenAlign != 0 {
		b.ApplyAlign()
	}
}
func (b *Box) CalculationDimension(parent Element) {
	solveWidth := b.Width
	solveHeight := b.Height
	if parent != nil && b.UserParentRelative {
		if b.ParentRelativeScaleWidth > 0 {
			solveWidth = parent.GetWidth() * b.ParentRelativeScaleWidth
		}
		if b.ParentRelativeScaleHeight > 0 {
			solveHeight = parent.GetHeight() * b.ParentRelativeScaleHeight
		}
		if b.ParentRelativeDeltaHeight != 0 && b.Height == 0 {
			solveHeight += b.ParentRelativeDeltaHeight
		}
		if b.ParentRelativeDeltaWidth != 0 && b.Width == 0 {
			solveWidth += b.ParentRelativeDeltaWidth
		}
	}

	b.DrawArea.Width = solveWidth
	b.DrawArea.Height = solveHeight

	// with padding
	b.Width = solveWidth + b.Padding.Left + b.Padding.Right
	b.Height = solveHeight + b.Padding.Top + b.Padding.Bottom

	// reassign with box sizing
	if b.BoxSizing == BorderBox {
		b.DrawArea.Width = b.DrawArea.Width - b.Padding.Left - b.Padding.Right
		b.DrawArea.Height = b.DrawArea.Height - b.Padding.Top - b.Padding.Bottom

		b.Width = b.Width - b.Padding.Left - b.Padding.Right
		b.Height = b.Height - b.Padding.Top - b.Padding.Bottom
	}

}
func (b *Box) ApplyChildPosition() {
	if b.Children != nil {
		// handle padding
		if b.LayoutManager != nil {
			b.LayoutManager.Place(b)
		}
		if b.ChildrenAlign != 0 {
			b.ApplyAlign()
		}
	}
}
func (b *Box) AddToChildren(element Element) {
	if b.Children == nil {
		b.Children = []Element{}
	}
	b.Children = append(b.Children, element)
}
func (b *Box) AlignWithAnchor(anchor Anchor, X float64, Y float64) {
	// transform anchor to topleft anchor
	originX := 0.0
	originY := 0.0
	if anchor == TopCenter {
		originY = b.Y
		originX = b.X - (b.Width / 2)
	}
	if anchor == TopRight {
		originY = b.Y
		originX = b.X - b.Width
	}
	if anchor == CenterLeft {
		originY = b.Y - (b.Height / 2)
		originX = b.X
	}
	if anchor == Center {
		originY = b.Y - (b.Height / 2)
		originX = b.X - (b.Width / 2)
	}
	if anchor == CenterRight {
		originY = b.Y - (b.Height / 2)
		originX = b.X - b.Width
	}
	if anchor == BottomLeft {
		originY = b.Y - b.Height
		originX = b.X
	}
	if anchor == BottomCenter {
		originY = b.Y - b.Height
		originX = b.X - (b.Width / 2)
	}
	if anchor == BottomRight {
		originY = b.Y - b.Height
		originX = b.X - b.Width
	}
	b.X = originX
	b.Y = originY
}

func (b *Box) GetChildren() []Element {
	return b.Children
}

func (b *Box) GetDimension() (float64, float64) {
	return b.Width, b.Height
}
func (b *Box) GetWeight() float64 {
	return b.Weight
}

func (b *Box) Render(renderContext *RenderContext) {
	context := renderContext.Context
	if len(b.Color) != 0 {
		context.DrawRectangle(b.X, b.Y, b.Width, b.Height)
		context.SetHexColor(b.Color)
		context.Fill()
	}
}

type Edge struct {
	Top    float64
	Right  float64
	Bottom float64
	Left   float64
}

type Anchor int

const (
	TopLeft Anchor = iota
	TopCenter
	TopRight
	CenterLeft
	Center
	CenterRight
	BottomLeft
	BottomCenter
	BottomRight
)

func (b *Box) GetAnchor(anchor Anchor) (X, Y float64) {
	if anchor == TopLeft {
		X = b.X
		Y = b.Y
	}
	if anchor == TopCenter {
		X = b.Width/2 + b.X
		Y = b.Y
	}
	if anchor == TopRight {
		X = b.X + b.Width
		Y = b.Y
	}
	if anchor == CenterLeft {
		X = b.X
		Y = b.Y + b.Height/2
	}
	if anchor == Center {
		X = b.Width/2 + b.X
		Y = b.Y + b.Height/2
	}
	if anchor == CenterRight {
		X = b.X + b.Width
		Y = b.Y + b.Height/2
	}
	if anchor == BottomLeft {
		X = b.X
		Y = b.Y + b.Height
	}
	if anchor == BottomCenter {
		X = b.Width/2 + b.X
		Y = b.Y + b.Height
	}
	if anchor == BottomRight {
		X = b.X + b.Width
		Y = b.Y + b.Height
	}
	return
}

type ChildAlign int

const (
	MainAxisTop ChildAlign = iota + 1
	MainAxisCenter
	MainAxisBottom
	CrossAxisLeft
	CrossAxisCenter
	CrossAxisRight
	MainCrossAxisCenter
)

func (b *Box) ApplyAlign() {
	childAlign := b.ChildrenAlign
	if childAlign == MainAxisTop {
		_, alignY := b.GetAnchor(TopLeft)
		for _, child := range b.Children {
			child.SetY(alignY)
		}
	}
	if childAlign == MainAxisCenter {
		_, alignY := b.GetAnchor(CenterLeft)
		for _, child := range b.Children {
			child.SetY(alignY - child.GetHeight()/2)
		}
	}
	if childAlign == MainAxisBottom {
		_, alignY := b.GetAnchor(BottomLeft)
		for _, child := range b.Children {
			child.SetY(alignY - child.GetHeight())
		}
	}
	if childAlign == CrossAxisLeft {
		alignX, _ := b.GetAnchor(BottomLeft)
		for _, child := range b.Children {
			child.SetX(alignX)
		}
	}
	if childAlign == CrossAxisCenter {
		alignX, _ := b.GetAnchor(BottomCenter)
		for _, child := range b.Children {
			child.SetX(alignX - child.GetWidth()/2)
		}
	}
	if childAlign == CrossAxisRight {
		alignX, _ := b.GetAnchor(BottomRight)
		for _, child := range b.Children {
			child.SetX(alignX - child.GetWidth())
		}
	}

	if childAlign == MainCrossAxisCenter {
		targetX, targetY := b.GetAnchor(Center)
		for _, child := range b.Children {
			child.SetX(targetX - child.GetWidth()/2)
			child.SetY(targetY - child.GetHeight()/2)
			fmt.Println(targetY - child.GetHeight()/2)
		}
	}
}
