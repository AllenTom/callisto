package callisto

type Dimension interface {
	GetDimension() (width float64, height float64)
	GetWidth() float64
	GetHeight() float64
	SetHeight(height float64)
	SetWidth(width float64)
	SetDimension(width float64, height float64)
	SetDimensionDelta(width float64, height float64)
	GetWeight() float64
	SetWeight(weight float64)
	CalculationDimension(parent Element)
}

type ElementDimension struct {
	Width                     float64
	Height                    float64
	Weight                    float64
	ParentRelativeScaleWidth  float64
	ParentRelativeScaleHeight float64
	ParentRelativeDeltaWidth  float64
	ParentRelativeDeltaHeight float64
	UserParentRelative        bool
}

func (e *ElementDimension) GetWeight() float64 {
	return e.Weight
}

func (e *ElementDimension) SetWeight(weight float64) {
	e.Weight = weight
}

func (e *ElementDimension) GetDimension() (float64, float64) {
	return e.Width, e.Height
}

func (e *ElementDimension) GetWidth() float64 {
	width, _ := e.GetDimension()
	return width
}

func (e *ElementDimension) GetHeight() float64 {
	_, height := e.GetDimension()
	return height
}

func (e *ElementDimension) SetHeight(height float64) {
	e.Height = height
}

func (e *ElementDimension) SetWidth(width float64) {
	e.Width = width
}

func (e *ElementDimension) SetDimension(width float64, height float64) {
	e.SetWidth(width)
	e.SetHeight(height)
}

func (e *ElementDimension) SetDimensionDelta(width float64, height float64) {
	e.Width += width
	e.Height += height
}
