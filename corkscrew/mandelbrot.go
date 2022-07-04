package corkscrew

/* Copyright © 2022 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

import "image"

type MandelbrotTile struct {
  X, Y   float64
  Bounds image.Rectangle
}

func NewMandelbrotTile(x, y float64, bounds image.Rectangle) *MandelbrotTile {
  m := &MandelbrotTile{
    X:      x,
    Y:      y,
    Bounds: bounds,
  }

  return m
}

func (m *MandelbrotTile) generate() {

}
