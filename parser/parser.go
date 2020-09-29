package parser

import (
	"encoding/json"
	"github.com/allentom/callisto"
	"io/ioutil"
)

type Meta struct {
	From         string  `json:"from"`
	TemplatePath string  `json:"templatePath"`
	Width        float64 `json:"width"`
	Height       float64 `json:"height"`
}
type Body struct {
	Children  []map[string]interface{} `json:"children"`
	Direction string                   `json:"direction"`
}
type Doc struct {
	Meta Meta `json:"meta"`
	Body Body `json:"body"`
}

func ReadJsonDoc(docPath string) (*Doc, error) {
	var err error
	var doc Doc
	raw, err := ioutil.ReadFile(docPath)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(raw, &doc)
	if err != nil {
		return nil, err
	}
	return &doc, nil
}

type BuildContext struct {
	parent  callisto.Element
	content map[string]interface{}
}

func BuildTree(engine *callisto.RenderEngine, sourceFilePath string) (err error) {
	doc, err := ReadJsonDoc(sourceFilePath)
	if err != nil {
		return
	}

	if doc.Meta.From == "empty" {
		err = engine.NewCanvas(int(doc.Meta.Width), int(doc.Meta.Height))
		if err != nil {
			return
		}
	}
	if doc.Meta.From == "template" {
		err = engine.LoadCanvasImageFromFile(doc.Meta.TemplatePath)
		if err != nil {
			return
		}
	}
	var layoutManager callisto.LayoutManager
	if doc.Body.Direction == "vertical" || doc.Body.Direction == "" {
		layoutManager = &callisto.VerticalListLayoutManager{}
	}
	if doc.Body.Direction == "horizon" {
		layoutManager = &callisto.HorizonListLayoutManager{}
	}
	root := &callisto.Box{
		ElementPosition: callisto.ElementPosition{
			UseParentDelta: true,
		},
		ElementDimension: callisto.ElementDimension{
			ParentRelativeScaleWidth:  1,
			ParentRelativeScaleHeight: 1,
			UserParentRelative:        true,
		},
		Children:      []callisto.Element{},
		LayoutManager: layoutManager,
	}

	buildQueue := make([]BuildContext, 0)
	for _, child := range doc.Body.Children {
		buildQueue = append(buildQueue, BuildContext{
			parent:  root,
			content: child,
		})
	}
	for {
		if len(buildQueue) == 0 {
			break
		}
		var context BuildContext
		if len(buildQueue) < 2 {
			context, buildQueue = buildQueue[0], make([]BuildContext, 0)
		} else {
			context, buildQueue = buildQueue[0], buildQueue[1:]
		}
		rawElementType, ok := context.content["element"]
		if !ok {
			continue
		}
		elementType, isString := rawElementType.(string)
		if !isString {
			continue
		}
		// parse node
		handler := Dispatch(elementType)
		if handler == nil {
			continue
		}
		node, err := handler(BuildContext{
			parent:  context.parent,
			content: context.content,
		})
		if err != nil {
			return err
		}
		context.parent.AddToChildren(node)
		//append child
		rawChildren, exist := context.content["children"]
		if !exist {
			continue
		}
		children, check := rawChildren.([]interface{})
		if !check {
			continue
		}

		for _, rawChild := range children {
			child := rawChild.(map[string]interface{})
			buildQueue = append(buildQueue, BuildContext{
				parent:  node,
				content: child,
			})
		}

	}
	engine.SetDoc(root)
	return
}

type ElementHandler func(context BuildContext) (callisto.Element, error)

func Dispatch(elementType string) ElementHandler {
	if elementType == "box" {
		return BoxHandler
	}
	if elementType == "text" {
		return TextHandler
	}
	if elementType == "img" {
		return ImageElementHandler
	}
	if elementType == "paragraph" {
		return ParagraphHandler
	}
	return nil
}
