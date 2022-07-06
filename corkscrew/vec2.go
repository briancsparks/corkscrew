package corkscrew

/* Copyright © 2022 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

type Vec2 struct {
  X, Y float32
}

func NewVec2(x float32, y float32) *Vec2 {
  return &Vec2{X: x, Y: y}
}

func V2(x, y float32) Vec2 {
  return Vec2{
    X: x,
    Y: y,
  }
}

func V2FromFloat64(x, y float64) Vec2 {
  return Vec2{float32(x), float32(y)}
}

func Dx(v1, v2 Vec2) float32 {
  return v2.X - v1.X
}

func Dy(v1, v2 Vec2) float32 {
  return v2.Y - v1.Y
}

func needNormalize(v1, v2 Vec2) bool {
  if v1.X < v2.X && v1.Y < v2.Y {
    return false
  }
  return true
}

func normalize(v1, v2 Vec2) (Vec2, Vec2) {
  if !needNormalize(v1, v2) {
    return v1, v2
  }

  x0 := mathMin(v1.X, v2.X)
  y0 := mathMin(v1.Y, v2.Y)
  x1 := mathMax(v1.X, v2.X)
  y1 := mathMax(v1.Y, v2.Y)

  return V2(x0, y0), V2(x1, y1)
}

func mathMin(a,b float32) float32 {
  if b < a {
    return b
  }
  return a
}

func mathMax(a,b float32) float32 {
  if a > b {
    return a
  }
  return b
}

func mathMinMax(a, b float32) (float32, float32) {
  if b < a {
    return b, a
  }
  return a, b
}

func mathCenterPoint(x1, x2 float32) float32 {
  return (x1 + x2) / 2.0
}

func minMax(x1, y1, x2, y2 float32, field *Field) (Vec2, Vec2) {
  return minMaxV2(V2(x1, y1), V2(x2, y2), field)
}

func minMaxV2(v1, v2 Vec2, field *Field) (Vec2, Vec2) {
  x1, x2 := mathMinMax(v1.X, v2.X)
  y1, y2 := mathMinMax(v1.Y, v2.Y)

  if field.ShowMathy {
    min := V2(x1, y2)
    max := V2(x2, y1)
    return min, max
  }

  min := V2(x1, y1)
  max := V2(x2, y2)
  return min, max
}

