package parser

import (
	"errors"
	"fmt"
	"github.com/allentom/callisto"
	"github.com/flopp/go-findfont"
	"strconv"
	"strings"
)

func GetLayoutManagerWithString(rawLayoutManager string) callisto.LayoutManager {
	switch rawLayoutManager {
	case "vertical":
		return &callisto.VerticalListLayoutManager{}
	case "horizon":
		return &callisto.HorizonListLayoutManager{}
	default:
		return &callisto.VerticalListLayoutManager{}
	}
}

var ChildAlignMapping map[string]callisto.ChildAlign = map[string]callisto.ChildAlign{
	"mainAxisTop":         callisto.MainAxisTop,
	"mainAxisCenter":      callisto.MainAxisCenter,
	"mainAxisBottom":      callisto.MainAxisBottom,
	"crossAxisLeft":       callisto.CrossAxisLeft,
	"crossAxisCenter":     callisto.CrossAxisCenter,
	"crossAxisRight":      callisto.CrossAxisRight,
	"crossMainAxisCenter": callisto.MainCrossAxisCenter,
}

func GetChildAlignWithString(rawChildAlign string) callisto.ChildAlign {
	align, isExist := ChildAlignMapping[rawChildAlign]
	if !isExist {
		return 0
	}
	return align
}

func GetPaddingFromRaw(rawString string) callisto.ElementPadding {
	elementPadding := callisto.ElementPadding{}
	paddingGroups := strings.Split(rawString, " ")
	for idx := range paddingGroups {
		value, err := strconv.ParseFloat(strings.ReplaceAll(paddingGroups[idx], "px", ""), 10)
		if err != nil {
			fmt.Println(err)
		}
		if idx == 0 {
			elementPadding.Top = value
		} else if idx == 1 {
			elementPadding.Right = value
		} else if idx == 2 {
			elementPadding.Bottom = value
		} else if idx == 3 {
			elementPadding.Left = value
		}
	}
	return elementPadding
}

func ParseBoxSizing(rawString string) callisto.BoxSizing {
	if rawString == "border-box" {
		return callisto.BorderBox
	}
	return callisto.ContentBox
}

var ImageFitMapping map[string]callisto.ImageFit = map[string]callisto.ImageFit{
	"width":  callisto.FitWidth,
	"height": callisto.FitHeight,
	"both":   callisto.Both,
	"none":   callisto.NoFit,
}

func ParseImageFit(rawString string) (callisto.ImageFit, error) {
	fit, exist := ImageFitMapping[rawString]
	if exist {
		return fit, nil
	}
	return 0, errors.New("image fit invalid")
}

func ParseFont(rawString string) string {
	if fileExists(rawString) {
		return rawString
	}
	fontNames := []string{rawString, "Arial"}
	for _, name := range fontNames {
		fontPath, err := findfont.Find(name)
		if err == nil {
			return fontPath
		}
	}
	return ""
}
