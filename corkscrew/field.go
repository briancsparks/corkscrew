package corkscrew

import "image"

/* Copyright © 2022 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

type Field struct {
  // Properties that are math-centric.
  // - They act like a math-head would expect (y increases UP the the screen) and such.
  Min, Max Vec2                                                     // (-50.0, -50.0), (50.0, 50.0)

  // Properties that are comptuer-centric
  Bounds image.Rectangle                                            // 0, 0, 800, 600

  // Display properties
  ShowXAxis, ShowYAxis, ShowGridLines bool
  IsLogX, IsLogY                      bool
  ShowMathy                           bool
}

func (f *Field) mkTile(w, h int, rw, rh, centerx, centery float32) *Tile {
  return NewTile(w, h , rw, rh, centerx, centery, f)
}


func (f *Field) Coordinate(t *Tile, pt image.Point) Vec2 {
  return Coordinate(t, pt, f.ShowMathy)
}

func (f *Field) FBounds(width, height float32) (Vec2, Vec2) {
  var min, max Vec2
  if f.ShowMathy {
    min = V2(-width / 2.0, height / 2.0)
    max = V2(width / 2.0, -height / 2.0)
  } else {
    min = V2(-width / 2.0, -height / 2.0)
    max = V2(width / 2.0, height / 2.0)
  }

  return min, max
}

func NewField(bounds image.Rectangle, width, height float32) *Field {     // (0,0,800,600), 100.0
  rx := width / 2.0                                                       // 50.0
  ry := height / 2.0                                                       // 50.0
  //brx := float32(bounds.Dx()) / 2.0                                       // 400.0    - bounds radius x
  //bry := float32(bounds.Dy()) / 2.0                                       // 300.0    - bounds radius y
  //fPerI := r / brx                                                        // 0.125    - units for the float part per int pixel
                                                                          //            Every int pixel is 0.125 float units wide

  f := &Field{
    Bounds:     bounds,
    //Min:        V2(-fPerI*brx, -fPerI*brx),
    //Max:        V2(fPerI*bry, fPerI*bry),
    Min:        V2(-rx, -ry),
    Max:        V2(rx, ry),
    ShowMathy:  true,
  }

  return f
}
