package callisto

import (
	"github.com/fogleman/gg"
)

type ImageBox struct {
	Box
	Src             string
	Image           *ImageElement
	OriginDimension bool
	Fit             ImageFit
}

type ImageFit uint

const (
	NoFit ImageFit = iota
	FitWidth
	FitHeight
	Both
)

func (b *ImageBox) CalculationDimension(parent Element) {
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

	b.Width = solveWidth
	b.Height = solveHeight
	if b.Fit == NoFit {
		b.Width = float64(b.Image.Config.Width)
		b.Height = float64(b.Image.Config.Height)
		return
	}
	if b.Fit == FitWidth {
		resizeScale := b.Width / float64(b.Image.Config.Width)
		b.Height = resizeScale * float64(b.Image.Config.Height)
	}
	if b.Fit == FitHeight {
		resizeScale := b.Height / float64(b.Image.Config.Height)
		b.Width = resizeScale * float64(b.Image.Config.Width)
	}
}
func (b *ImageBox) Preload(ctx *RenderContext) {
	ctx.ImageResourceLibrary.LoadFromResource(b.Src)
	b.Image = ctx.ImageResourceLibrary.GetImageWithPath(b.Src)

}
func (b *ImageBox) Render(renderContext *RenderContext) {
	context := renderContext.Context
	imageContext := gg.NewContext(int(b.Width), int(b.Height))
	imageContext.ScaleAbout(b.Width/float64(b.Image.Config.Width), b.Height/float64(b.Image.Config.Height), 0, 0)
	imageContext.DrawImage(*b.Image.Content, 0, 0)
	context.DrawImage(imageContext.Image(), int(b.X), int(b.Y))
}
