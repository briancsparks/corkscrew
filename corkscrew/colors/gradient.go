package colors

/* Copyright © 2022 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

import (
	"github.com/lucasb-eyer/go-colorful"
	"os"
	"strconv"
)

// -------------------------------------------------------------------------------------------------------------------

func GetGetColor() func(iterations, maxIterations int) colorful.Color {

  palette := gradientGen()
  black, _ := colorful.Hex("#000000")

  return func(iterations, maxIterations int) colorful.Color {
    percentage := float32(iterations) / float32(maxIterations)
    index := int(percentage * float32(len(palette)))

    var result colorful.Color

    if index >= len(palette) {
      result = black
    } else {
      result = palette[index]
    }

    //fmt.Printf("[%06d] %v idx: %v Picked color: %v\n", iterations, percentage, index, result)
    return result
  }

}

// This table contains the "keypoints" of the colorgradient you want to generate.
// The position of each keypoint has to live in the range [0,1]

type GradientTable []struct {
	Col colorful.Color
	Pos float64
}

// This is the meat of the gradient computation. It returns a HCL-blend between
// the two colors around `t`.
// Note: It relies heavily on the fact that the gradient keypoints are sorted.

func (gt GradientTable) GetInterpolatedColorFor(t float64) colorful.Color {
	for i := 0; i < len(gt)-1; i++ {
		c1 := gt[i]
		c2 := gt[i+1]
		if c1.Pos <= t && t <= c2.Pos {
			// We are in between c1 and c2. Go blend them!
			t := (t - c1.Pos) / (c2.Pos - c1.Pos)
			return c1.Col.BlendHcl(c2.Col, t).Clamped()
		}
	}

	// Nothing found? Means we're at (or past) the last gradient keypoint.
	return gt[len(gt)-1].Col
}

func MustParseHex(s string) colorful.Color {
	c, err := colorful.Hex(s)
	if err != nil {
		panic("MustParseHex: " + err.Error())
	}
	return c
}

func gradientGen() []colorful.Color {
	// The "keypoints" of the gradient.
	keypoints := GradientTable{
		{MustParseHex("#9e0142"), 0.0},
		{MustParseHex("#d53e4f"), 0.1},
		{MustParseHex("#f46d43"), 0.2},
		{MustParseHex("#fdae61"), 0.3},
		{MustParseHex("#fee090"), 0.4},
		{MustParseHex("#ffffbf"), 0.5},
		{MustParseHex("#e6f598"), 0.6},
		{MustParseHex("#abdda4"), 0.7},
		{MustParseHex("#66c2a5"), 0.8},
		{MustParseHex("#3288bd"), 0.9},
		{MustParseHex("#5e4fa2"), 1.0},
	}

	h := 1024

	if len(os.Args) == 3 {
		// Meh, I'm being lazy...
		h, _ = strconv.Atoi(os.Args[2])
	}

	var result []colorful.Color

	for y := h - 1; y >= 0; y-- {
		c := keypoints.GetInterpolatedColorFor(float64(y) / float64(h))
		result = append(result, c)
	}

	return result
}
