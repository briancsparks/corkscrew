package corkscrew

import (
  "github.com/lucasb-eyer/go-colorful"
  "image"
  "image/color"
  "image/draw"
)

/* Copyright © 2022 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */


func init() {

}

type Tile struct {
  Img      *image.RGBA
  Min, Max Vec2
  Rect     image.Rectangle

  sub      *SubTilizer
}

func NewTile(w, h int, rw, rh, centerx, centery float32, field *Field) *Tile {
  rect := image.Rect(0, 0, w, h)
  min, max := minMax(centerx-rw, centery-rh, centerx+rw, centery+rh, field)

  t := &Tile{
    Min:  min,
    Max:  max,
    Rect: rect,
    Img:  image.NewRGBA(rect),
  }
  t.sub = NewSubTilizer(t, field)

  c := colorful.WarmColor()
  img := &image.Uniform{C: color.RGBA{R: uint8(c.R * 255), G: uint8(c.G * 255), B: uint8(c.B * 255), A: 255}}
  draw.Draw(t.Img, t.Img.Bounds(), img, image.Point{}, draw.Src)

  return t
}

type tileMaker interface {
  mkTile(w, h int, rw, rh, centerx, centery float32) *Tile
}
