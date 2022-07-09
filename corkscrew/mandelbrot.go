package corkscrew

/* Copyright © 2022 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

import (
  "fmt"
  "image"
  "image/draw"
  "math"
  "sync"
  "time"
)

type MandelbrotTile struct {
  ID      int32
  Min, Max Vec2
  Bounds image.Rectangle

  joe *Joe
  d           DebugIt
}

// TODO: remove Joe

func NewMandelbrotTile(id int32, min, max Vec2/*, x, y float64*/, bounds image.Rectangle, joe *Joe) *MandelbrotTile {
  Validate(min, max)
  m := &MandelbrotTile{
    ID:     id,
    Min:    min,
    Max:    max,
    Bounds: bounds,
    d:      *NewDebugIt(id, "Mandelbrot"),
  }
  m.joe = joe
  m.d.Disable()

  return m
}

func (m *MandelbrotTile) Run(quit chan struct{}, tileMaker tileMaker, tilechan chan *Tile) error {

  wg := sync.WaitGroup{}

  wg.Add(1)
  go func() {
    wg.Done()
    printfBegin := m.d.BuildTaggedPrintf("begin")
    printfAllPoints := m.d.BuildTaggedPrintf("allPoints")
    printfEnd := m.d.BuildTaggedPrintf("end")

    start := time.Now()

    totalLoops := 0
    totalPixels := 0
    totalFails := 0
    totalFast := 0

    //z := 0+0i
    z := complex(float32(0.0), float32(0.0))
    c := complex(float32(0.0), float32(0.0))

    maxLoopWorkCount := 150
    loopWorkCount := 0
    iterations := 0
    maxIterations := 100
    maxDist := 2.0
    thresholdDist := float32(maxDist * maxDist)
    distSq := float32(0.0)

    //v := V2(0.0, 0.0)
    //z := complex(v.X, v.Y)

    var pixel image.Point
    var v2 Vec2

    fwidth := Dx(m.Min, m.Max)
    fheight := Dy(m.Min, m.Max)
    centerx := mathCenterPoint(m.Min.X, m.Max.X)
    centery := mathCenterPoint(m.Min.Y, m.Max.Y)
    tile := tileMaker.mkTile(m.ID, m.Bounds.Min.X, m.Bounds.Min.Y, m.Bounds.Max.X, m.Bounds.Max.Y, fwidth, fheight, centerx, centery)
    //pixelCount := tile.sub.PixelCount()
    iterCounts := make([]int, maxIterations + 2)

    tileDone := false
    fetchNext := false
    fetchFirst := true

    // If we need to fetch another to work on...
    if fetchFirst {
      // Get initial
      pv2, ppixel := tile.sub.Curr()
      if pv2 == nil || ppixel == nil {
        tileDone = true
      }
      v2, pixel = *pv2, *ppixel

      c = complex(v2.X, v2.Y)
      iterations = 0
      totalPixels += 1
      printfBegin("Beginning work on pixel (%4d, %4d) [point: (%7.5f, %7.5f)]\n", pixel.X, pixel.Y, v2.X, v2.Y)
    }

    if !tileDone {
      for {
        select {
        case <-quit:
          break

        default:
          iterations++
          iterations--
          break
        }

        var tooFar, tooMany bool
        var doneWithCurrent bool

        // If we need to fetch another to work on...
        if fetchNext {
          // Get initial
          pv2, ppixel := tile.sub.Next()
          if pv2 == nil || ppixel == nil {
            tileDone = true
            break
          }
          v2, pixel = *pv2, *ppixel

          z = complex(float32(0.0), float32(0.0))
          c = complex(v2.X, v2.Y)
          iterations = 0
          totalPixels += 1
          printfBegin("Beginning work on pixel (%4d, %4d) [point: (%7.5f, %7.5f)]\n", pixel.X, pixel.Y, v2.X, v2.Y)
        }

        //index := tile.sub.GetCurrIndex()
        for loopWorkCount = 0; loopWorkCount < maxLoopWorkCount; loopWorkCount++ {
          z = z*z + c
          iterations++
          totalLoops += 1

          assert(!math.IsNaN(float64(real(z))) && !math.IsNaN(float64(imag(z))))
          //assertMsg(iterations < 4, fmt.Sprintf("Iterations!: %v", iterations))

          if iterations > maxIterations {
            tooMany = true
            totalFails += 1
            break
          }

          // How far are we?
          re, im := real(z), imag(z)
          distSq = re*re + im*im
          if distSq > thresholdDist {
            // We are off in the weeds
            tooFar = true
            totalFast += 1
            break
          }
          //assertMsg(distSq <= thresholdDist, fmt.Sprintf("distSq: %v, (%v)", distSq, thresholdDist))
        }
        printfAllPoints("%05d -- workCount: (%v/%v), Iterated (%v/%v), pix: %8v, c: %11v,  z: %11v,  d: %9v, %v\n",
             totalLoops, loopWorkCount+1, maxLoopWorkCount, iterations, maxIterations, pixel, c, z, distSq, status(tooMany, tooFar, tileDone))

        if tooMany {
          // Its stuck
          doneWithCurrent = true
          fetchNext = true
        } else if tooFar {
          // Too far away, will never come back
          doneWithCurrent = true
          fetchNext = true
        } else {
          // Ran out of time for this run. Not an error, just loop around
        }

        // If we are done with this pixel, report it or whatever
        if doneWithCurrent {
          iterCounts[iterations] += 1
          if pixel.X == m.Bounds.Min.X {
            tilechan <- tile
          }

          color := getColor(iterations, maxIterations)
          //draw.Draw(tile.Img, tile.Img.Bounds(), &image.Uniform{color}, image.Point{}, draw.Src)
          //draw.Draw(tile.Img, image.Rectangle{image.Point{0, 0}, pixel}, &image.Uniform{color}, image.Point{}, draw.Src)

          onePixRect := image.Rect(pixel.X, pixel.Y, pixel.X + 1, pixel.Y + 1)
          draw.Draw(tile.Img, onePixRect, &image.Uniform{C: color}, image.Point{}, draw.Src)

          printfEnd("Final #%v pixel: %v  Z: %v, C: %v, %v\n", iterations, color, z, c, status(tooMany, tooFar, tileDone))
          joe.dataChannels.messages <- fmt.Sprintf("[%2d]: %v: %v: %v %v", iterations, pixel, z, color)
        }

      }
    }

    if tileDone {
      elapsed := time.Since(start)
      m.d.PrintfReport("Time: %f sec -- Loops: [%v], Pixels: %v, fast: %v, fail: %v\n", elapsed.Seconds(), totalLoops, totalPixels, totalFast, totalFails)
      m.d.PrintfReport("Iter counts: %v\n", iterCounts)
      tilechan <- tile
    }
  }()
  wg.Wait()

  return nil
}

func status(tooMany, tooFar, tileDone bool) string {
  s := ""

  if tooMany {
    s += "tooMany"
  }

  if tooFar {
    s = appendTo(s, "tooFar")
  }

  s = "[" + s + "]"

  if tileDone {
    s = appendsTo(s, " ", "Done")
  }
  return s
}
