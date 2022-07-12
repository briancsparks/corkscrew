package corkscrew

/* Copyright © 2022 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */
import (
  "github.com/briancsparks/corkscrew/corkscrew/both"
  "github.com/go-p5/p5"
)

// -------------------------------------------------------------------------------------------------------------------

func (c *MandelBothCmd) DrawP5() {
  func() {
    c.lock.Lock()
    defer c.lock.Unlock()

    p5.DrawImage(c.mandelImg, 0, 0)
  }()

  c.showAxes()
  c.showGridLines()
}

// --------------------------------------------------------------------------------------------------------------------

func (c *MandelBothCmd) showAxes() {
  var x,y float64

  x = 0.0
  y = 0.0
  origin := c.grid.DisplayPtFor(both.WorkPt{X: x, Y: y})

  // The origin and major x-y axis
  p5.StrokeWidth(2)
  p5.Line(float64(c.grid.Main.Display.Min.X), float64(origin.Y), float64(c.grid.Main.Display.Max.X), float64(origin.Y))
  p5.Line(float64(origin.X), float64(c.grid.Main.Display.Min.Y), float64(origin.X), float64(c.grid.Main.Display.Max.Y))
}

// --------------------------------------------------------------------------------------------------------------------

func (c *MandelBothCmd) showGridLines() {
  var x, y, dx, dy float64

  x = 0.0
  y = 0.0

  // The origin and major x-y axis
  p5.StrokeWidth(2)

  // Grid lines (x-->)
  dx/*, dy*/ = 0.0/*, 0.0*/
  for dx = 1.0; dx < c.grid.Main.Work.Max.X; dx += 1.0 {
    gridPt := c.grid.DisplayPtFor(both.WorkPt{X: x+dx, Y: y})
    p5.Line(float64(gridPt.X), float64(c.grid.Main.Display.Min.Y), float64(gridPt.X), float64(c.grid.Main.Display.Max.Y))
  }

  // Grid lines (<--x)
  dx/*, dy*/ = 0.0/*, 0.0*/
  for dx = -1.0; dx >= c.grid.Main.Work.Min.X; dx -= 1.0 {
    gridPt := c.grid.DisplayPtFor(both.WorkPt{X: x+dx, Y: y})
    p5.Line(float64(gridPt.X), float64(c.grid.Main.Display.Min.Y), float64(gridPt.X), float64(c.grid.Main.Display.Max.Y))
  }

  // Grid lines (y-->)
  dx, dy = 0.0, 0.0
  for dy = 1.0; dy < c.grid.Main.Work.Min.Y; dy += 1.0 {
    gridPt := c.grid.DisplayPtFor(both.WorkPt{X: x, Y: y+dy})
    p5.Line(float64(c.grid.Main.Display.Min.X), float64(gridPt.Y), float64(c.grid.Main.Display.Max.X), float64(gridPt.Y))
  }

  // Grid lines (<--y)
  dx, dy = 0.0, 0.0
  for dy = -1.0; dy >= c.grid.Main.Work.Max.Y; dy -= 1.0 {
   gridPt := c.grid.DisplayPtFor(both.WorkPt{X: x, Y: y+dy})
   p5.Line(float64(c.grid.Main.Display.Min.X), float64(gridPt.Y), float64(c.grid.Main.Display.Max.X), float64(gridPt.Y))
  }
}

