package corkscrew

/* Copyright © 2022 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

import (
  "github.com/go-p5/p5"
)

type Joe struct {
  Field *Field
  Tiles []*Tile
}

func NewJoe(f *Field) *Joe {
  j := &Joe{
    Field: f,
  }

  j.Tiles = append(j.Tiles, NewTile(70, 70, 5.0, 5.0, 0.0))

  return j
}

func (j *Joe) Run() {

}

func (j *Joe) Render() {

  for _, tile := range j.Tiles {
    p5.DrawImage(tile.Img, float64(tile.Min.X), float64(tile.Min.Y))
  }
}

