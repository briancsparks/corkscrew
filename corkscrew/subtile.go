package corkscrew

/* Copyright © 2022 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

import "image"


func init() {

}

type SubTilizer struct {
  Field           *Field
  ParentTile      *Tile

  CurrentIndexPoint image.Point
  CurrentIndex    int

  // Computed and stored
  currVec2  Vec2
  currPoint image.Point
}

func NewSubTilizer(parent *Tile, field *Field) *SubTilizer {
  return &SubTilizer{
    Field:        field,
    ParentTile:   parent,
  }
}

//func (st *SubTilizer) NextX() (Vec2, image.Point) {
//  index := st.CurrentIndex + 1
//  p := st.ParentTile
//  pt := Coord2d(&p.Rect, index)
//  x, y := Coordinate4(p.Min, p.Max, &p.Rect, pt, st.Field.ShowMathy)
//
//  st.currVec2, st.currPoint, st.CurrentIndex = Vec2{x,y}, pt, index
//
//  return st.currVec2, st.currPoint
//}

func (st *SubTilizer) Next() (Vec2, image.Point) {
  v, pt, index := st.at(st.CurrentIndex + 1)
  st.currVec2, st.currPoint, st.CurrentIndex = v, pt, index

  return st.currVec2, st.currPoint
}

func (st *SubTilizer) Current() (Vec2, image.Point) {
  return st.currVec2, st.currPoint
}

func (st *SubTilizer) at(index int) (Vec2, image.Point, int) {
  p := st.ParentTile
  pt := Coord2d(&p.Rect, index)
  x, y := Coordinate4(p.Min, p.Max, &p.Rect, pt, st.Field.ShowMathy)

  return Vec2{x,y}, pt, index
}

func (st *SubTilizer) validIndex(index int) bool {
  r := st.ParentTile.Rect
  area := r.Dx() * r.Dy()

  return index > 0 && index < area
}

func Coordinate(t *Tile, pt image.Point, mathy bool) Vec2 {
  x, y := Coordinate4(t.Min, t.Max, &t.Rect, pt, mathy)
  return V2(x, y)
}

// Coordinate4 computes the center of a pixel as a float
func Coordinate4(min, max Vec2, rect *image.Rectangle, pt image.Point, mathy bool) (float32, float32) {
  fx := halfCoord(min.X, max.X, rect.Min.X, rect.Max.X, rect.Dx(), pt.X, false)
  fy := halfCoord(min.Y, max.Y, rect.Min.Y, rect.Max.Y, rect.Dx(), pt.Y, mathy)

  return fx, fy
}

func halfCoord(fmin, fmax float32, imin, imax int, width int, index int, backwards bool) float32 {

  cellwidth, half, offset := computeHalfHelper(fmin, fmax, imin, imax, index)

  var centerPoint float32
  if  backwards {
    centerPoint = fmax - half - (float32(offset) * cellwidth)
  } else {
    centerPoint = fmin + half + (float32(offset) * cellwidth)
  }

  return centerPoint
}

func computeHalfHelper(fmin, fmax float32, imin, imax int, index int) ( cellwidth, half float32, offset int) {
  fwidth := fmax - fmin
  iwidth := imax - imin
  cellwidth = fwidth / float32(iwidth)        // An int cell is this much wide in the float numbers
  half = cellwidth / 2.0
  offset = index - imin

  return
}

func Coord2d(r *image.Rectangle, index int) image.Point {
  x := index % r.Dx()
  y := index / r.Dy()

  return image.Pt(x, y)
}

