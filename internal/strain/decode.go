package strain

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/2785/n471-proj-carrot/model"
)

const (
	gax string = "gax"
	gay string = "gay"
	gaz string = "gaz"
	asx string = "asx"
	asy string = "asy"
	asz string = "asz"
	xx  string = "xx"
	xy  string = "xy"
	xz  string = "xz"
	yx  string = "yx"
	yy  string = "yy"
	yz  string = "yz"
	zx  string = "zx"
	zy  string = "zy"
	zz  string = "zz"
)

var re = regexp.MustCompile(`Ga (?P<gax>-?\d+\.?\d*(?:e?-?\d+)?) (?P<gay>-?\d+\.?\d*(?:e?-?\d+)?) (?P<gaz>-?\d+\.?\d*(?:e?-?\d+)?)\nAs (?P<asx>-?\d+\.?\d*(?:e?-?\d+)?) (?P<asy>-?\d+\.?\d*(?:e?-?\d+)?) (?P<asz>-?\d+\.?\d*(?:e?-?\d+)?)\n---\nX refaxis: \[(?P<xx>-?\d+\.?\d*(?:e?-?\d+)?), (?P<xy>-?\d+\.?\d*(?:e?-?\d+)?), (?P<xz>-?\d+\.?\d*(?:e?-?\d+)?)\]\nY ref axis \[(?P<yx>-?\d+\.?\d*(?:e?-?\d+)?), (?P<yy>-?\d+\.?\d*(?:e?-?\d+)?), (?P<yz>-?\d+\.?\d*(?:e?-?\d+)?)\]\nZ refaxis \[(?P<zx>-?\d+\.?\d*(?:e?-?\d+)?), (?P<zy>-?\d+\.?\d*(?:e?-?\d+)?), (?P<zz>-?\d+\.?\d*(?:e?-?\d+)?)\]`)

// Decode .
func Decode(s string) (*model.Input, error) {
	match := re.FindStringSubmatch(s)

	if match == nil {
		return nil, fmt.Errorf("no match")
	}

	result := make(map[string]string)
	for i, name := range re.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = match[i]
		}
	}

	floats := make(map[string]float64)

	for _, v := range []string{
		gax,
		gay,
		gaz,
		asx,
		asy,
		asz,
		xx,
		xy,
		xz,
		yx,
		yy,
		yz,
		zx,
		zy,
		zz} {
		val, ok := result[v]
		if !ok || val == "" {
			return nil, fmt.Errorf("%s not found", v)
		}

		f, err := strconv.ParseFloat(val, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse %s into float: %w", val, err)
		}

		floats[v] = f
	}

	return &model.Input{
		Ga:   model.Coord{X: floats[gax], Y: floats[gay], Z: floats[gaz]},
		As:   model.Coord{X: floats[asx], Y: floats[asy], Z: floats[asz]},
		XRef: model.Coord{X: floats[xx], Y: floats[xy], Z: floats[xz]},
		YRef: model.Coord{X: floats[yx], Y: floats[yy], Z: floats[yz]},
		ZRef: model.Coord{X: floats[zx], Y: floats[zy], Z: floats[zz]},
	}, nil

}
