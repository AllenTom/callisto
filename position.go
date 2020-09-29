package callisto

type Position interface {
	GetPosition() (float64, float64)
	SetPosition(float64, float64)
	UpdatePosition(float64, float64) (float64, float64)
	MoveDelta(float64, float64)
	GetX() float64
	GetY() float64
	SetY(float64)
	SetX(float64)
	TransformChildPosition()
}

type ElementPosition struct {
	X               float64
	Y               float64
	ParentDeltaX    float64
	ParentDeltaY    float64
	SiblingDeltaX   float64
	SiblingDeltaY   float64
	UseParentDelta  bool
	UseSiblingDelta bool
	siblingAnchor   Anchor
}

func (b *ElementPosition) SetY(y float64) {
	b.Y = y
}

func (b *ElementPosition) SetX(x float64) {
	b.X = x
}

func (b *ElementPosition) GetX() float64 {
	return b.X
}

func (b *ElementPosition) GetY() float64 {
	return b.Y
}

func (b *ElementPosition) GetPosition() (float64, float64) {
	return b.X, b.Y
}

func (b *ElementPosition) MoveDelta(x float64, y float64) {
	b.X += x
	b.Y += y
}
func (b *ElementPosition) SetPosition(x float64, y float64) {
	b.X = x
	b.Y = y
}
func (b *ElementPosition) UpdatePosition(x float64, y float64) (float64, float64) {
	b.X = x
	b.Y = y
	return b.X, b.Y
}
