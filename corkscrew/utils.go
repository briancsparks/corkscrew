package corkscrew

/* Copyright © 2022 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

import "github.com/lucasb-eyer/go-colorful"

var palette []colorful.Color

func init() {
  palette = gradientGen()
  //fmt.Printf("palette: %v\n", palette)
  //fmt.Printf("len: %v\n", len(palette))
}

func getColor(iterations int) colorful.Color {
  if iterations > len(palette) {
    return palette[len(palette) - 1]
  }
  return palette[iterations]
}
