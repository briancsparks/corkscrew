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

  black, _ = colorful.Hex("#000000")
  white, _ = colorful.Hex("#000000")
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

func asserter(test bool) bool {
  if !test {
    breakout("", true)
  }
  return !test
}

func assertMsg(test bool, msg string) {
  if !test {
    breakout(msg, false)
  }
}

func assert(test bool) {
  if !test {
    breakout("", false)
  }
}

func breakout(msg string, quiet bool) {
  if !quiet {
    fmt.Printf("  ------------ BREAKOUT!! %v !!\n", msg)
  }
}
