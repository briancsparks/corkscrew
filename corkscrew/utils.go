package corkscrew

/* Copyright © 2022 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

import (
  "fmt"
  colorful "github.com/lucasb-eyer/go-colorful"
)

var palette []colorful.Color
var black, white colorful.Color

func init() {
  palette = gradientGen()
  //black = colorful.Color{R: 0.313725, G: 0.478431, B: 0.721569}
  //black = colorful.Color{R: 0.0, G: 0.0, B: 0.0}

  black, _ = colorful.Hex("#000000")

  //fmt.Printf("palette: %v\n", palette)
  fmt.Printf("len: %v\n", len(palette))
}

func getColor(iterations, maxIterations int) colorful.Color {
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

func assertMsg(test bool, msg string) {
  if !test {
    breakout(msg)
  }
}

func assert(test bool) {
  if !test {
    breakout("")
  }
}

func breakout(msg string) {
  fmt.Printf("  ------------ BREAKOUT!! %v !!\n", msg)
}
