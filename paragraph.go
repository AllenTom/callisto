package callisto

import (
	"fmt"
	"github.com/fogleman/gg"
)

type ParagraphBox struct {
	Box
	Text     string
	FontPath string
	Size     float64
	Color    string
	Wrap     bool
}

func (t *ParagraphBox) Preload(ctx *RenderContext) {
	drawContext := ctx.Context
	err := drawContext.LoadFontFace(t.FontPath, t.Size)
	if err != nil {
		fmt.Println(err)
	}
	if !t.Wrap {
		return
	}
	width, height := drawContext.MeasureString(t.Text)
	if width == 0 {
		t.Width = width
	}
	if t.Height == 0 {
		t.Height = height
	}
}

func (t *ParagraphBox) Render(renderContext *RenderContext) {
	drawContext := renderContext.Context
	err := drawContext.LoadFontFace(t.FontPath, t.Size)
	if err != nil {
		fmt.Println(err)
	}
	drawContext.SetHexColor(t.Color)
	drawContext.WordWrap(t.Text, t.Width)
	drawContext.DrawStringWrapped(t.Text, t.X, t.Y, 0, 0, t.Width, 2, gg.AlignLeft)
}
