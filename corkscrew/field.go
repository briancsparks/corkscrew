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
  Min, Max Vec2                                                     // (-50.0, 50.0), (50.0, -50.0)

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

func (f *Field) mkTile(id int, l, t, r, b int, rw, rh, centerx, centery float32) *Tile {
  return NewTile(id, l, t, r, b, rw, rh, centerx, centery, f)
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

  return f
}

// --------------------------------------------------------------------------------------------------------------------

func (f *Field) getRealmPixel(x, y float32) *image.Point {
  //func RealmToScreen(pt Vec2, realmRect Rec2, rec image.Rectangle) image.Point {
  pt := RealmToScreen(Vec2{x, y}, Rec2{f.Min, f.Max}, f.Bounds)
  return &pt
}

// --------------------------------------------------------------------------------------------------------------------

func (f *Field) RenderBg() {

}

// --------------------------------------------------------------------------------------------------------------------

func (f *Field) Render() {
  //count++
  //t := time.Now()
  //elapsed := t.Sub(startTime).Seconds()
  //sec := t.Second()
  //
  //clockStart := -(math.Pi / 2)
  //
  //p5.StrokeWidth(2)
  //p5.Fill(color.RGBA{R: 255, A: 208})
  //p5.Ellipse(50, 50, 80, 80)
  //
  //p5.Fill(color.RGBA{B: 255, A: 208})
  //p5.Quad(50, 50, 80, 50, 80, 120, 60, 120)
  //
  //p5.Fill(color.RGBA{G: 255, A: 208})
  //p5.Rect(200, 200, 50, 100)
  //
  //p5.Fill(color.RGBA{G: 255, A: 208})
  //p5.Triangle(100, 100, 120, 120, 80, 120)
  //
  //p5.TextSize(24)
  //p5.Text(fmt.Sprintf("%d - %v", count, float64(count)/elapsed), 10, 300)
  //
  //p5.Stroke(color.Black)
  //p5.StrokeWidth(5)
  //p5.Arc(300, 100, 80, 80, clockStart, clockStart+(float64(sec)/60.0)*2.0*math.Pi)

}

// --------------------------------------------------------------------------------------------------------------------

func (f *Field) showHorizAxis() {
  var x,y float32

  x = 0.0
  y = 0.0
  origin := f.getRealmPixel(x, y)

  // The origin and major x-y axis
  p5.StrokeWidth(2)
  //p5.Line(float64(origin.X), float64(f.Bounds.Min.Y), float64(origin.X), float64(f.Bounds.Max.Y))
  p5.Line(float64(f.Bounds.Min.X), float64(origin.Y), float64(f.Bounds.Max.X), float64(origin.Y))
}

// --------------------------------------------------------------------------------------------------------------------

func (f *Field) showVertAxis() {
  var x,y float32

  x = 0.0
  y = 0.0
  origin := f.getRealmPixel(x, y)

  // The origin and major x-y axis
  p5.StrokeWidth(2)
  p5.Line(float64(origin.X), float64(f.Bounds.Min.Y), float64(origin.X), float64(f.Bounds.Max.Y))
  //p5.Line(float64(f.Bounds.Min.X), float64(origin.Y), float64(f.Bounds.Max.X), float64(origin.Y))
}

// --------------------------------------------------------------------------------------------------------------------

func (f *Field) showGridLines() {
  var x,y,dx,dy float32

  x = 0.0
  y = 0.0
  //origin := f.getRealmPixel(x, y)

  // The origin and major x-y axis
  p5.StrokeWidth(2)
  //p5.Line(float64(origin.X), float64(f.Bounds.Min.Y), float64(origin.X), float64(f.Bounds.Max.Y))
  //p5.Line(float64(f.Bounds.Min.X), float64(origin.Y), float64(f.Bounds.Max.X), float64(origin.Y))

  // Grid lines (x-->)
  dx, dy = 0.0, 0.0
  for dx = 1.0; dx < f.Max.X; dx += 1.0 {
    gridPt := f.getRealmPixel(x + dx, y)
    p5.Line(float64(gridPt.X), float64(f.Bounds.Min.Y), float64(gridPt.X), float64(f.Bounds.Max.Y))
  }

  // Grid lines (<--x)
  dx, dy = 0.0, 0.0
  for dx = -1.0; dx >= f.Min.X; dx -= 1.0 {
    gridPt := f.getRealmPixel(x + dx, y)
    p5.Line(float64(gridPt.X), float64(f.Bounds.Min.Y), float64(gridPt.X), float64(f.Bounds.Max.Y))
  }

  // Grid lines (y-->)
  dx, dy = 0.0, 0.0
  for dy = 1.0; dy < f.Min.Y; dy += 1.0 {
    gridPt := f.getRealmPixel(x, y + dy)
    p5.Line(float64(f.Bounds.Min.X), float64(gridPt.Y), float64(f.Bounds.Max.X), float64(gridPt.Y))
  }

  // Grid lines (<--y)
  dx, dy = 0.0, 0.0
  for dy = -1.0; dy >= f.Max.Y; dy -= 1.0 {
    gridPt := f.getRealmPixel(x, y + dy)
    p5.Line(float64(f.Bounds.Min.X), float64(gridPt.Y), float64(f.Bounds.Max.X), float64(gridPt.Y))
  }
}

// --------------------------------------------------------------------------------------------------------------------

func (f *Field) RenderLast() {

  //configOptions.ShowHorizAxis = false
  if configOptions.ShowHorizAxis {
    f.showHorizAxis()
  }

  //configOptions.ShowVertAxis = false
  if configOptions.ShowVertAxis {
    f.showVertAxis()
  }

  //configOptions.ShowGridLines = false
  if configOptions.ShowGridLines {
    f.showHorizAxis()
    f.showVertAxis()
    f.showGridLines()
  }

}



