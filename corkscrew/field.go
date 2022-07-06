package corkscrew

/* Copyright © 2022 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

import (
  "fmt"
  "github.com/go-p5/p5"
  "image"
  "math"
)

// --------------------------------------------------------------------------------------------------------------------

type Field struct {
  // Properties that are math-centric.
  // - They act like a math-head would expect (y increases UP the screen) and such.
  Min, Max Vec2                                                     // (-50.0, -50.0), (50.0, 50.0)

  // Properties that are comptuer-centric
  Bounds image.Rectangle                                            // 0, 0, 800, 600

  // Display properties
  //ShowXAxis, ShowYAxis, ShowGridLines bool
  IsLogX, IsLogY                      bool
  ShowMathy                           bool
}

// --------------------------------------------------------------------------------------------------------------------

func (f *Field) UnitsPerPix() (Vec2, Vec2) {
  iMinX     := f.Bounds.Min.X
  iMinY     := f.Bounds.Min.Y

  iMaxX     := f.Bounds.Max.X
  iMaxY     := f.Bounds.Max.Y

  return UnitsPerPix6(f.Min, f.Max, iMinX, iMinY, iMaxX, iMaxY)
}

// --------------------------------------------------------------------------------------------------------------------

func UnitsPerPix6(fMin, fMax Vec2, iMinX, iMinY, iMaxX, iMaxY int) (Vec2, Vec2) {
  fMinX     := fMin.X
  fMinY     := fMin.Y

  fMaxX     := fMax.X
  fMaxY     := fMax.Y

  return UnitsPerPix8(fMinX, fMinY, fMaxX, fMaxY, iMinX, iMinY, iMaxX, iMaxY)
}

// --------------------------------------------------------------------------------------------------------------------

func UnitsPerPix8(fMinX, fMinY, fMaxX, fMaxY float32, iMinX, iMinY, iMaxX, iMaxY int) (Vec2, Vec2) {
  fwidth    := float32(math.Abs(float64(fMaxX - fMinX)))
  fheight   := float32(math.Abs(float64(fMaxY - fMinY)))

  iwidth    := math.Floor(math.Abs(float64(iMaxX - iMinX)))
  iheight   := math.Floor(math.Abs(float64(iMaxY - iMinY)))

  unitsPPX  := fwidth / float32(iwidth)
  pixelsPUX := float32(iwidth) / fwidth

  unitsPPY  := fheight / float32(iheight)
  pixelsPUY := float32(iheight) / fheight

  return Vec2{unitsPPX, unitsPPY}, Vec2{pixelsPUX, pixelsPUY}
}

// --------------------------------------------------------------------------------------------------------------------

func (f *Field) mkTile(w, h int, rw, rh, centerx, centery float32) *Tile {
  return NewTile(w, h , rw, rh, centerx, centery, f)
}

// --------------------------------------------------------------------------------------------------------------------

func (f *Field) Coordinate(t *Tile, pt image.Point) Vec2 {
  return Coordinate(t, pt, f.ShowMathy)
}

// --------------------------------------------------------------------------------------------------------------------

func (f *Field) FBounds() (Vec2, Vec2) {
  return f.Min, f.Max
}

// --------------------------------------------------------------------------------------------------------------------

func NewField(bounds image.Rectangle, /*width, height*/ left, top, right, bottom float32) *Field {     // (0,0,800,600), 100.0

  f := &Field{
    Bounds:     bounds,
    ShowMathy:  true,
  }
  min, max := minMax(left, top, right, bottom, f)

  f.Min = min
  f.Max = max

  // Some stats
  unitsPerPixel, pixelsPerUnit := f.UnitsPerPix()
  fmt.Printf("UnitsPerPixel: %v, PixelsPerUnit %v\n", unitsPerPixel, pixelsPerUnit)

  // Show the config
  fmt.Printf("Show grid lines: %v\n", configOptions.ShowGridLines)
  fmt.Printf("Show horizontal axis: %v\n", configOptions.ShowHorizAxis)
  fmt.Printf("Show vertical axis: %v\n", configOptions.ShowVertAxis)

  //v2Look := Vec2{0, 0}
  //_, origin := GridPt2(v2Look, f.Min, f.Max, f.Bounds, f.ShowMathy)
  //fmt.Printf("GridPt. Pt: %v, f: %v %v, r: %v\n", v2Look, f.Min, f.Max, f.Bounds)
  //fmt.Printf("GridLines: origin: %v; bounds: %v\n", origin, f.Bounds)

  return f
}

// --------------------------------------------------------------------------------------------------------------------

func (f *Field) getPixel(x, y float32) *image.Point {
  _, pt := GridPt2(Vec2{x, y}, f.Min, f.Max, f.Bounds, f.ShowMathy)
  return pt
}

// --------------------------------------------------------------------------------------------------------------------

func (f *Field) Render() {

}

// --------------------------------------------------------------------------------------------------------------------

func (f *Field) RenderLast() {

  if configOptions.ShowGridLines {
    var x,y/*,dx,dy*/ float32

    x = 0.0
    y = 0.0
    origin := f.getPixel(x, y)
    //v2Look := Vec2{x, y}
    //_, origin := GridPt2(v2Look, f.Min, f.Max, f.Bounds, f.ShowMathy)

    // The origin and major x-y axis
    p5.StrokeWidth(2)
    p5.Line(float64(origin.X), float64(f.Bounds.Min.Y), float64(origin.X), float64(f.Bounds.Max.Y))
    p5.Line(float64(f.Bounds.Min.X), float64(origin.Y), float64(f.Bounds.Max.X), float64(origin.Y))

    //// Grid lines (x-->)
    //dx, dy = 0.0, 0.0
    //for dx = 1.0; dx < f.Max.X; dx += 1.0 {
    //  gridPt := f.getPixel(x + dx, y)
    //  p5.Line(float64(gridPt.X), float64(f.Bounds.Min.Y), float64(gridPt.X), float64(f.Bounds.Max.Y))
    //}
    //
    //// Grid lines (<--x)
    //dx, dy = 0.0, 0.0
    //for dx = -1.0; dx >= f.Min.X; dx -= 1.0 {
    //  gridPt := f.getPixel(x + dx, y)
    //  p5.Line(float64(gridPt.X), float64(f.Bounds.Min.Y), float64(gridPt.X), float64(f.Bounds.Max.Y))
    //}
    //
    //// Grid lines (y-->)
    //dx, dy = 0.0, 0.0
    //for dy = 1.0; dy < f.Max.Y; dy += 1.0 {
    //  gridPt := f.getPixel(x, y + dy)
    //  p5.Line(float64(f.Bounds.Min.X), float64(gridPt.Y), float64(f.Bounds.Max.X), float64(gridPt.Y))
    //}
    //
    //// Grid lines (<--y)
    //dx, dy = 0.0, 0.0
    //for dy = -1.0; dy >= f.Min.Y; dy -= 1.0 {
    //  gridPt := f.getPixel(x, y + dy)
    //  p5.Line(float64(f.Bounds.Min.X), float64(gridPt.Y), float64(f.Bounds.Max.X), float64(gridPt.Y))
    //}

  }

  if configOptions.ShowHorizAxis {

  }

  if configOptions.ShowVertAxis {

  }

}



