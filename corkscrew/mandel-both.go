package corkscrew

import (
  "fmt"
  "github.com/briancsparks/corkscrew/corkscrew/both"
  "github.com/lucasb-eyer/go-colorful"
  "image"
  "image/color"
  "image/draw"
  "math"
  "sync"
  "time"
)

/* Copyright © 2022 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

// ------------------------------------------------------------------------------------------------------------------

type MandelDataMessage struct {
  Id        int
  Img      *image.RGBA
  Rect      image.Rectangle
}

// ------------------------------------------------------------------------------------------------------------------

func RunMandel(quit chan struct{}, tilechan chan *MandelDataMessage, params *MandelConfig, bothGrid *both.Both) error {
  fmt.Printf("params: %v\n", params)

  c := colorful.WarmColor()
  imageUniform := &image.Uniform{C: color.RGBA{R: uint8(c.R * 255), G: uint8(c.G * 255), B: uint8(c.B * 255), A: 255}}

  grid1, grid2 := bothGrid.Split()
  gridA, gridB := grid1.Split()
  cmd.grids = append(cmd.grids, gridA, gridB)
  gridA, gridB = grid2.Split()
  cmd.grids = append(cmd.grids, gridA, gridB)
  //grids := []*both.Both{grid1, grid2}

  for _, grid := range cmd.grids {

    wg := sync.WaitGroup{}
    wg.Add(1)
    go func(grid *both.Both) {
      wg.Done()
      start := time.Now()

      r := grid.Main.Display
      displayRect := image.Rect(r.Min.X, r.Min.Y, r.Max.X, r.Max.Y)

      fmt.Printf("grid: %v\n", *grid)
      computeCount := 0
      iterations, maxIterations := 0, 100

      // Stats to collect
      totalIterations := 0
      totalPixels := 0
      totalFails := 0
      totalFast := 0
      iterCounts := make([]int, maxIterations + 2)

      var pixel both.DisplayPt

      pixImage := image.NewRGBA(displayRect)
      draw.Draw(pixImage, displayRect.Bounds(), imageUniform, image.Point{}, draw.Src)
      tilechan <- &MandelDataMessage{grid.Id, pixImage, displayRect}

      var bothPt *both.BothPt
      itPt := grid.GetDisplayIterator()
      bothPt = itPt.Curr()
      pixel = bothPt.Display
      totalPixels += 1

      z := complex(0.0, 0.0)
      c := complex(bothPt.Work.X, bothPt.Work.Y)

      doneWithCurrent := false
      allDone := false
      for {
        select {
        case <-quit:
          break

        default:
          break
        }

        if bothPt == nil {
          bothPt = itPt.Next()
          if bothPt == nil {
            // Done
            allDone = true
            break
          }
          z = complex(0.0, 0.0)
          c = complex(bothPt.Work.X, bothPt.Work.Y)

          // BBB
          if computeCount >= 540000 + 580 {
            computeCount += 1
          }
          computeCount += 1
          totalPixels += 1
          pixel = bothPt.Display
          doneWithCurrent = false
        }

        for iterations = 1; iterations < maxIterations; iterations++ {
          z = z*z + c

          // How far away are we?
          re, im := real(z), imag(z)
          distSq := re*re + im*im
          if distSq > 4.0 {
            doneWithCurrent = true
            totalFast += 1
            break
          }
          if math.IsNaN(re) || math.IsNaN(im) || math.IsNaN(distSq) {
            break
          }
        }
        totalIterations += iterations
        if iterations >= maxIterations {
          totalFails += 1
        }

        if doneWithCurrent || iterations >= maxIterations {
          iterCounts[iterations] += 1
          if pixel.X == grid.Main.Display.Min.X {
            tilechan <- &MandelDataMessage{grid.Id, pixImage, displayRect}
          }

          kolor := getColor(iterations, maxIterations)
          onePixRect := image.Rect(pixel.X, pixel.Y, pixel.X + 1, pixel.Y + 1)
          draw.Draw(pixImage, onePixRect, &image.Uniform{C: kolor}, image.Point{}, draw.Src)

          bothPt = nil
        }
      }

      if allDone {
        elapsed := time.Since(start)
        fmt.Printf("Time: %f sec -- Loops: [%v], Pixels: %v, fast: %v, fail: %v\n", elapsed.Seconds(), totalIterations, totalPixels, totalFast, totalFails)
        fmt.Printf("Iter counts: %v\n", iterCounts)
        tilechan <- &MandelDataMessage{1, pixImage, displayRect}
      }
    }(grid)
    wg.Wait()
  }


  return nil
}








