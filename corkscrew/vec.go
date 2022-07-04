package corkscrew

/* Copyright © 2022 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

type Vec2 struct {
  X, Y float32
}

func V2(x, y float32) Vec2 {
  return Vec2{
    X: x,
    Y: y,
  }
}
