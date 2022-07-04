package corkscrew

import "image"

/* Copyright © 2022 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */


func init() {

}

func Coordinate(t *Tile, pt *image.Point, mathy bool) Vec2 {
  x, y := Coordinate4(&t.Min, &t.Max, &t.Rect, pt, mathy)
  return V2(x, y)
}

// Coordinate4 computes the center of a pixel as a float
func Coordinate4(min, max *Vec2, rect *image.Rectangle, pt *image.Point, mathy bool) (float32, float32) {
  fx := halfCoord(min.X, max.X, rect.Min.X, rect.Max.X, pt.X, false)
  fy := halfCoord(min.Y, max.Y, rect.Min.Y, rect.Max.Y, pt.Y, mathy)

  return fx, fy
}

func halfCoord(fmin, fmax float32, imin, imax int, x int, backwards bool) float32 {
  fwidth := fmax - fmin
  iwidth := imax - imin
  icellwidth := fwidth / float32(iwidth)
  halficellwidth := icellwidth / 2.0
  index := x - imin

  var centerPoint float32
  if  backwards {
    centerPoint = fmax - halficellwidth - (float32(index) * icellwidth)
  } else {
    centerPoint = fmin + halficellwidth + (float32(index) * icellwidth)
  }

  return centerPoint
}


