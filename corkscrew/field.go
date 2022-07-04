package corkscrew

import "image"

/* Copyright © 2022 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

type Field struct {
  // Properties that are math-centric.
  // - They act like a math-head would expect (y increases UP the the screen) and such.
  Min, Max Vec2

  // Properties that are comptuer-centric
  Bounds image.Rectangle

  // Display properties
  ShowXAxis, ShowYAxis, ShowGridLines bool
  IsLogX, IsLogY                      bool
}

func NewField(bounds image.Rectangle, width float32) *Field {
  r := width / 2.0
  brx := float32(bounds.Dx()) / 2.0
  bry := float32(bounds.Dy()) / 2.0
  fPerI := r / brx

  f := &Field{
    Bounds: bounds,
    Min:    V2(-fPerI*brx, -fPerI*brx),
    Max:    V2(fPerI*bry, fPerI*bry),
  }

  return f
}
