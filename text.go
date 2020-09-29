package callisto

import "fmt"

type TextBox struct {
	Box
	Text     string
	FontPath string
	Size     float64
	Color    string
}

func (t *TextBox) Preload(ctx *RenderContext) {
	drawContext := ctx.Context
	err := drawContext.LoadFontFace(t.FontPath, t.Size)
	if err != nil {
		fmt.Println(err)
	}
	t.Width, t.Height = drawContext.MeasureString(t.Text)
}

func (t TextBox) Render(renderContext *RenderContext) {
	drawContext := renderContext.Context
	err := drawContext.LoadFontFace(t.FontPath, t.Size)
	if err != nil {
		fmt.Println(err)
	}
	drawContext.SetHexColor(t.Color)
	drawContext.DrawStringAnchored(t.Text, t.DrawArea.X, t.DrawArea.Y, 0, 1)
}
