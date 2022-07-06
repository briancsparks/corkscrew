package corkscrew

/* Copyright © 2022 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

type Rec2 struct {
  Min     Vec2
  Max     Vec2
}

func NewRec2(min Vec2, max Vec2) *Rec2 {
  return &Rec2{Min: min, Max: max}
}

func R2(x0,y0,x1,y1 float32) Rec2 {
  return Rec2{Min: Vec2{x0,y0}, Max: Vec2{x1,y1}}
}

func (r *Rec2) Parts() (float32, float32, float32, float32) {
  return r.Min.X, r.Min.Y, r.Max.X, r.Max.Y
}

func (r *Rec2) Dx() float32 {
  return r.Max.X - r.Min.X
}

func (r *Rec2) Dy() float32 {
  diff := r.Max.Y - r.Min.Y
  if diff > 0 {
    return diff
  }
  return -diff
}
