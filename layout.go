package callisto

type LayoutManager interface {
	Place(box *Box)
}
type VerticalListLayoutManager struct {
}

func (v *VerticalListLayoutManager) Place(element *Box) {
	offsetY := element.DrawArea.Y
	for _, child := range element.GetChildren() {
		child.SetX(element.DrawArea.X)
		child.SetY(offsetY)
		offsetY += child.GetHeight()
	}
}

type HorizonListLayoutManager struct {
}

func (v *HorizonListLayoutManager) Place(element *Box) {
	offsetX := element.DrawArea.X
	for _, child := range element.GetChildren() {
		child.SetY(element.DrawArea.Y)
		child.SetX(offsetX)
		offsetX += child.GetWidth()
	}
}
