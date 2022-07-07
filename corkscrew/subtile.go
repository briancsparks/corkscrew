package corkscrew

/* Copyright © 2022 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

import (
  "image"
)

type SubTilizer struct {
  Field           *Field
  ParentTile      *Tile

  CurrentIndexPoint image.Point
}

func NewSubTilizer(parent *Tile, field *Field) *SubTilizer {
  return &SubTilizer{
    Field:        field,
    ParentTile:   parent,
  }
}

func (st *SubTilizer) Curr() (*Vec2, *image.Point) {
  p := st.ParentTile
  pt := st.CurrentIndexPoint

  index := (pt.Y * p.Rect.Dx()) + pt.X
  if index > p.Rect.Dx() * p.Rect.Dy() {
    return nil, nil
  }

  x := halfCoord(p.Min.X, p.Max.X, p.Rect.Min.X, p.Rect.Max.X, pt.X, false)
  y := halfCoord(p.Min.Y, p.Max.Y, p.Rect.Min.Y, p.Rect.Max.Y, pt.Y, false)

  st.CurrentIndexPoint = pt

  return &Vec2{x, y}, &pt
}

func (st *SubTilizer) Next() (*Vec2, *image.Point) {
  p := st.ParentTile
  pt := st.CurrentIndexPoint

  pt.X += 1
  if pt.X >= p.Rect.Dx() {
    pt.X = 0
    pt.Y += 1
  }

  //index := pt.X * pt.Y
  index := (pt.Y * p.Rect.Dx()) + pt.X
  if index >= p.Rect.Dx() * p.Rect.Dy() {
    return nil, nil
  }

  x := halfCoord(p.Min.X, p.Max.X, p.Rect.Min.X, p.Rect.Max.X, pt.X, false)
  y := halfCoord(p.Min.Y, p.Max.Y, p.Rect.Min.Y, p.Rect.Max.Y, pt.Y, false)

  st.CurrentIndexPoint = pt

  return &Vec2{x, y}, &pt
}

func halfCoord(fmin, fmax float32, imin, imax int, midpt int, backwards bool) float32 {
  widthOfIcell, half, offset := computeHalfHelper(fmin, fmax, imin, imax, midpt)

  var center float32
  if backwards {
   center = fmax - half - (float32(offset) * widthOfIcell)
  } else {
   center = fmin + half + (float32(offset) * widthOfIcell)
  }

  return center
}

func computeHalfHelper(fmin, fmax float32, imin, imax int, index int) ( cellwidth, half float32, offset int) {
  fwidth := fmax - fmin
  iwidth := imax - imin
  cellwidth = fwidth / float32(iwidth)        // An int cell is this much wide in the float numbers
  half = cellwidth / 2.0
  offset = index - imin

  return
}

