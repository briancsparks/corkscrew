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
}

func NewTile(w, h int, rw, rh, center float32) *Tile {
  rect := image.Rect(0, 0, w, h)

  t := &Tile{
    Min:  V2(center-rw, center+rw),
    Max:  V2(center-rh, center+rh),
    Rect: rect,
    Img:  image.NewRGBA(rect),
  }

  c := colorful.WarmColor()
  img := &image.Uniform{C: color.RGBA{R: uint8(c.R * 255), G: uint8(c.G * 255), B: uint8(c.B * 255), A: 255}}
  draw.Draw(t.Img, t.Img.Bounds(), img, image.Point{}, draw.Src)

  return t
}
