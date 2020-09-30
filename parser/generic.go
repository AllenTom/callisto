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

func ParseBorder(rawString string) (callisto.ElementBorder, error) {
	border := callisto.ElementBorder{
		Top:    callisto.Border{},
		Right:  callisto.Border{},
		Bottom: callisto.Border{},
		Left:   callisto.Border{},
	}
	// not set border
	if len(rawString) == 0 {
		return border, nil
	}
	width, color, err := ParseBorderWithPosition(rawString)
	if err != nil {
		return border, err
	}

	border.Top.Color = color
	border.Right.Color = color
	border.Bottom.Color = color
	border.Left.Color = color

	border.Top.Width = width
	border.Right.Width = width
	border.Bottom.Width = width
	border.Left.Width = width

	return border, nil
}

func ParseBorderWithPosition(rawString string) (width int, color string, err error) {
	if len(rawString) == 0 {
		return
	}
	rawGroup := strings.Split(rawString, " ")
	if len(rawGroup) != 2 {
		err = errors.New("unexpected border value")
		return
	}
	for _, rawValue := range rawGroup {
		if strings.Contains(rawValue, "#") {
			color = rawValue
		}

		if strings.Contains(rawValue, "px") {
			parseWidth, err := strconv.Atoi(strings.ReplaceAll(rawValue, "px", ""))
			if err != nil {
				err = errors.New("unexpected border width value")
				return 0, "", err
			}
			width = parseWidth
		}
	}
	return
}
