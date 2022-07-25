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

    //p5.DrawImage(c.mandelImg, 0, 0)
    for i, img := range c.mandelImgs {
      rect := c.mandelRects[i]
      p5.DrawImage(img, float64(rect.Min.X), float64(rect.Min.Y))
    }
  }()

  if Config.ShowAxis {
    c.showAxes()
  }

  if Config.ShowGridLines {
    c.showGridLines()
  }
}

// --------------------------------------------------------------------------------------------------------------------

func (c *MandelBothCmd) showAxes() {

  x, y := 0.0, 0.0
  origin := c.grid.DisplayPtFor(both.WorkPt{X: x, Y: y})

  // The origin and major x-y axis
  p5.StrokeWidth(2)
  p5.Line(float64(c.grid.Main.Display.Min.X), float64(origin.Y), float64(c.grid.Main.Display.Max.X), float64(origin.Y))
  p5.Line(float64(origin.X), float64(c.grid.Main.Display.Min.Y), float64(origin.X), float64(c.grid.Main.Display.Max.Y))

}

// --------------------------------------------------------------------------------------------------------------------

func (c *MandelBothCmd) showGridLines() {
  //boundsAreSame := true
  //if c.previousDrawWorkBounds == nil || *c.previousDrawWorkBounds != c.grid.Main.Work {
  //  boundsAreSame = false
  //  c.previousDrawWorkBounds = &c.grid.Main.Work
  //}

  // FIXME: the labels need to have the same accuracy. Run to see: go run main/corkscrew.go mandelboth -x -0.16 -y 1.0405 -r 0.1

  p5.StrokeWidth(1)
  plotableX := FindBestLineForLabel(c.grid.Main.Work.Min.X, c.grid.Main.Work.Max.X)
  plotableY := FindBestLineForLabel(c.grid.Main.Work.Max.Y, c.grid.Main.Work.Min.Y)
  //if !boundsAreSame {
  //  fmt.Printf("bounds: %v\n", c.grid.Main.Work)
  //  fmt.Printf("plotable x: %v (%v, %v)\n", plotableX, c.grid.Main.Work.Min.X, c.grid.Main.Work.Max.X)
  //  fmt.Printf("plotable y: %v (%v, %v)\n", plotableY, c.grid.Main.Work.Max.Y, c.grid.Main.Work.Min.Y)
  //}

  for _, x := range plotableX {
    displayPt := c.grid.DisplayPtFor(both.WorkPt{x, plotableY[0]})
    displayX/*, displayY*/ := displayPt.X/*, displayPt.Y*/
    //if !boundsAreSame {
    //  fmt.Printf("plotable x: %v, y:%v\n", displayX, displayPt.Y)
    //}

    //p5.Line(float64(c.grid.Main.Display.Min.X), float64(displayY), float64(c.grid.Main.Display.Max.X), float64(displayY))
    p5.Line(float64(displayX), float64(c.grid.Main.Display.Min.Y), float64(displayX), float64(c.grid.Main.Display.Max.Y))
  }

  for _, y := range plotableY {
    displayPt := c.grid.DisplayPtFor(both.WorkPt{plotableX[0], y})
    /*displayX,*/ displayY := /*displayPt.X,*/ displayPt.Y
    //if !boundsAreSame {
    //  fmt.Printf("plotable x: %v, y:%v\n", displayPt.X, displayY)
    //}

    p5.Line(float64(c.grid.Main.Display.Min.X), float64(displayY), float64(c.grid.Main.Display.Max.X), float64(displayY))
    //p5.Line(float64(displayX), float64(c.grid.Main.Display.Min.Y), float64(displayX), float64(c.grid.Main.Display.Max.Y))
  }

}

// --------------------------------------------------------------------------------------------------------------------

//func (c *MandelBothCmd) showGridLines() {
//
//  x, y, dx, dy := 0.0, 0.0, 0.0, 0.0
//
//  // The origin and major x-y axis
//  p5.StrokeWidth(2)
//
//  // Grid lines (x-->)
//  dx, dy = 0.0, 0.0
//  for dx = 1.0; dx < c.grid.Main.Work.Max.X; dx += 1.0 {
//    gridPt := c.grid.DisplayPtFor(both.WorkPt{X: x+dx, Y: y})
//    p5.Line(float64(gridPt.X), float64(c.grid.Main.Display.Min.Y), float64(gridPt.X), float64(c.grid.Main.Display.Max.Y))
//  }
//
//  // Grid lines (<--x)
//  dx, dy = 0.0, 0.0
//  for dx = -1.0; dx >= c.grid.Main.Work.Min.X; dx -= 1.0 {
//    gridPt := c.grid.DisplayPtFor(both.WorkPt{X: x+dx, Y: y})
//    p5.Line(float64(gridPt.X), float64(c.grid.Main.Display.Min.Y), float64(gridPt.X), float64(c.grid.Main.Display.Max.Y))
//  }
//
//  // Grid lines (y-->)
//  dx, dy = 0.0, 0.0
//  for dy = 1.0; dy < c.grid.Main.Work.Min.Y; dy += 1.0 {
//    gridPt := c.grid.DisplayPtFor(both.WorkPt{X: x, Y: y+dy})
//    p5.Line(float64(c.grid.Main.Display.Min.X), float64(gridPt.Y), float64(c.grid.Main.Display.Max.X), float64(gridPt.Y))
//  }
//
//  // Grid lines (<--y)
//  dx, dy = 0.0, 0.0
//  for dy = -1.0; dy >= c.grid.Main.Work.Max.Y; dy -= 1.0 {
//   gridPt := c.grid.DisplayPtFor(both.WorkPt{X: x, Y: y+dy})
//   p5.Line(float64(c.grid.Main.Display.Min.X), float64(gridPt.Y), float64(c.grid.Main.Display.Max.X), float64(gridPt.Y))
//  }
//}

