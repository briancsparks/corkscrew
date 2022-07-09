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

// -------------------------------------------------------------------------------------------------------------------

type Tile struct {
  ID        int

  Img      *image.RGBA
  Min, Max Vec2
  Rect     image.Rectangle

  sub      *SubTilizer
}

// -------------------------------------------------------------------------------------------------------------------

func NewTile(id int, l, t, r, b int, rw, rh, centerx, centery float32, field *Field) *Tile {
  rrx, rry := rw / 2.0, rh / 2.0
  min, max := minMax(centerx-rrx, centery-rry, centerx+rrx, centery+rry, field)
  rect := image.Rect(l, t, r, b)

  tile := &Tile{
    ID:   id,
    Min:  min,
    Max:  max,
    Rect: rect,
    Img:  image.NewRGBA(rect),
  }
  tile.sub = NewSubTilizer(tile, field)

  c := colorful.WarmColor()
  img := &image.Uniform{C: color.RGBA{R: uint8(c.R * 255), G: uint8(c.G * 255), B: uint8(c.B * 255), A: 255}}
  draw.Draw(tile.Img, tile.Img.Bounds(), img, image.Point{}, draw.Src)

  return tile
}

// -------------------------------------------------------------------------------------------------------------------

type tileMaker interface {
  mkTile(id int, l, t, r, b int, rw, rh, centerx, centery float32) *Tile
}
