package corkscrew

import "image"

/* Copyright © 2022 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

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

  index := pt.X * pt.Y
  if index > p.Rect.Dx() * p.Rect.Dy() {
    return nil, nil
  }

  x := st.halfCoord(p.Min.X, p.Max.X, p.Rect.Min.X, p.Rect.Max.X, pt.X, false)
  y := st.halfCoord(p.Min.Y, p.Max.Y, p.Rect.Min.Y, p.Rect.Max.Y, pt.Y, st.Field.ShowMathy)

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

  index := pt.X * pt.Y
  if index > p.Rect.Dx() * p.Rect.Dy() {
    return nil, nil
  }

  x := st.halfCoord(p.Min.X, p.Max.X, p.Rect.Min.X, p.Rect.Max.X, pt.X, false)
  y := st.halfCoord(p.Min.Y, p.Max.Y, p.Rect.Min.Y, p.Rect.Max.Y, pt.Y, st.Field.ShowMathy)

  st.CurrentIndexPoint = pt

  return &Vec2{x, y}, &pt
}

func (st *SubTilizer) halfCoord(fmin, fmax float32, imin, imax int, midpt int, backwards bool) float32 {
  widthOfIcell, half, offset := computeHalfHelper(fmin, fmax, imin, imax, midpt)

  var center float32
  if backwards {
    center = fmax - half - (float32(offset) * widthOfIcell)
  } else {
    center = fmin + half + (float32(offset) * widthOfIcell)
  }

  return center
}

func (st *SubTilizer) Coord1d(fmin, fmax float32, imin, imax, midpt int) (widthOfIcell, half float32, offset int) {
  fwidth := fmax - fmin
  iwidth := imax - imin
  widthOfIcell = fwidth / float32(iwidth)
  half = widthOfIcell / 2.0
  offset = midpt - imin

  return widthOfIcell, half, offset
}


