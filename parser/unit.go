package parser

import (
	"strconv"
	"strings"
)

type Unit int

const (
	ParentScaleRelative Unit = iota + 1
	ParentDeltaRelative
	Pixel
)

var UnitSymbolMapping map[Unit]string = map[Unit]string{
	ParentDeltaRelative: "rd",
	ParentScaleRelative: "%",
	Pixel:               "px",
}

func ParseValueAndUnit(raw string) (float64, Unit, error) {
	var err error
	var valueUnit Unit
	for unit, symbol := range UnitSymbolMapping {
		if strings.Contains(raw, symbol) {
			valueUnit = unit
			raw = strings.ReplaceAll(raw, symbol, "")
			break
		}
	}

	value := 0.0
	value, err = strconv.ParseFloat(raw, 10)
	if err != nil {
		return 0, 0, err
	}

	return value, valueUnit, nil
}
