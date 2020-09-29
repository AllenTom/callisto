package callisto

import "github.com/fogleman/gg"

type RenderEngine struct {
	Context *gg.Context
	Root    Element
}

func (e *RenderEngine) LoadCanvasImageFromFile(imageFilePath string) error {
	template, err := ReadImageFromFile(imageFilePath)
	e.Context = gg.NewContextForImage(*template)
	if err != nil {
		return err
	}
	return nil
}

func (e *RenderEngine) NewCanvas(width int, height int) error {
	e.Context = gg.NewContext(width, height)
	return nil
}

func (e *RenderEngine) SetDoc(doc Element) {
	if e.Context == nil {
		return
	}
	e.Root = &Box{
		ElementPosition: ElementPosition{},
		ElementDimension: ElementDimension{
			Width:  float64(e.Context.Width()),
			Height: float64(e.Context.Height()),
		},
		Children: []Element{
			doc,
		},
	}
}

func (e *RenderEngine) CalculationElementsDimension(imageResources *ImageResourceLibrary) {
	drawQueue := make([]*RenderQueueElement, 0)
	drawQueue = append(drawQueue, &RenderQueueElement{
		Root: e.Root,
		Context: &RenderContext{
			Context:              e.Context,
			ImageResourceLibrary: imageResources,
			Parent:               nil,
		},
	})
	for {
		var node *RenderQueueElement
		if len(drawQueue) == 1 {
			node, drawQueue = drawQueue[0], []*RenderQueueElement{}
		} else {
			node, drawQueue = drawQueue[len(drawQueue)-1], drawQueue[:len(drawQueue)-1]
		}
		for _, child := range node.Root.GetChildren() {
			preloadElement, needPreload := child.(PreloadElement)
			if needPreload {
				preloadElement.Preload(&RenderContext{
					Context:              e.Context,
					ImageResourceLibrary: imageResources,
					Parent:               node.Root,
					Sibling:              nil,
				})
			}
		}
		node.Root.CalculationDimension(node.Context.Parent)
		for _, childContext := range node.Root.GetChildren() {
			drawQueue = append(drawQueue, &RenderQueueElement{
				Root: childContext,
				Context: &RenderContext{
					Context:              e.Context,
					ImageResourceLibrary: imageResources,
					Parent:               node.Root,
				},
			})
		}
		if len(drawQueue) == 0 {
			break
		}
	}
}

func (e *RenderEngine) PlaceElements(imageResources *ImageResourceLibrary) {
	drawQueue := make([]*RenderQueueElement, 0)
	drawQueue = append(drawQueue, &RenderQueueElement{
		Root: e.Root,
		Context: &RenderContext{
			Context:              e.Context,
			ImageResourceLibrary: imageResources,
			Parent:               nil,
		},
	})
	for {
		var node *RenderQueueElement
		if len(drawQueue) == 1 {
			node, drawQueue = drawQueue[0], []*RenderQueueElement{}
		} else {
			node, drawQueue = drawQueue[0], drawQueue[1:]
		}

		node.Root.TransformChildPosition()
		for _, childContext := range node.Root.GetChildren() {
			drawQueue = append(drawQueue, &RenderQueueElement{
				Root: childContext,
				Context: &RenderContext{
					Context:              e.Context,
					ImageResourceLibrary: imageResources,
					Parent:               node.Root,
				},
			})
		}
		if len(drawQueue) == 0 {
			break
		}
	}
}

func (e *RenderEngine) RenderElements(imageResources *ImageResourceLibrary) {
	drawQueue := make([]*RenderQueueElement, 0)
	drawQueue = append(drawQueue, &RenderQueueElement{
		Root: e.Root,
		Context: &RenderContext{
			Context:              e.Context,
			ImageResourceLibrary: imageResources,
			Parent:               nil,
		},
	})
	for {
		var node *RenderQueueElement
		if len(drawQueue) == 1 {
			node, drawQueue = drawQueue[0], []*RenderQueueElement{}
		} else {
			node, drawQueue = drawQueue[0], drawQueue[1:]
		}
		node.Root.Render(node.Context)
		for _, childContext := range node.Root.GetChildren() {
			drawQueue = append(drawQueue, &RenderQueueElement{
				Root: childContext,
				Context: &RenderContext{
					Context:              e.Context,
					ImageResourceLibrary: imageResources,
					Parent:               node.Root,
				},
			})
		}
		if len(drawQueue) == 0 {
			break
		}
	}
}

func (e *RenderEngine) RenderImage() {
	imageResources := ImageResourceLibrary{
		Images: []*ImageElement{},
	}
	e.CalculationElementsDimension(&imageResources)
	e.PlaceElements(&imageResources)
	e.RenderElements(&imageResources)
}

func (e *RenderEngine) SaveAsPNG(outputImageFilePath string) error {
	err := e.Context.SavePNG(outputImageFilePath)
	if err != nil {
		return err
	}
	return nil
}
