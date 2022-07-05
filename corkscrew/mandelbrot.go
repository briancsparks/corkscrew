package corkscrew

/* Copyright © 2022 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

import (
  "fmt"
  "image"
  "image/draw"
  "math"
  "sync"
)

type MandelbrotTile struct {
  //X, Y   float64
  Min, Max Vec2
  Bounds image.Rectangle

  joe *Joe
}

// TODO: remove Joe

func NewMandelbrotTile(min, max Vec2/*, x, y float64*/, bounds image.Rectangle, joe *Joe) *MandelbrotTile {
  m := &MandelbrotTile{
    //X:      x,
    //Y:      y,
    Min:    min,
    Max:    max,
    Bounds: bounds,
  }
  m.joe = joe

  return m
}

func (m *MandelbrotTile) Run(quit chan struct{}, tileMaker tileMaker, tilechan chan *Tile) error {

  wg := sync.WaitGroup{}

  wg.Add(1)
  go func() {
    wg.Done()

    totalLoops := 0
    totalPixels := 0
    totalFails := 0
    totalFast := 0

    //z := 0+0i
    z := complex(float32(0.0), float32(0.0))
    c := complex(float32(0.0), float32(0.0))

    maxCount := 150
    count := 0
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
    tile := tileMaker.mkTile(m.Bounds.Dx(), m.Bounds.Dy(), fwidth, fheight, centerx, centery)

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
      //fmt.Printf("--Putting0 #%v pixel: %v  z: %v c: %v, [[tooMany: %v, tooFar: %v, done: %v]]\n", iterations, pixel, z, c, "N/A", "N/A", tileDone)
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
          //fmt.Printf("--Putting1 #%v pixel: %v  z: %v c: %v, [[tooMany: %v, tooFar: %v, done: %v]]\n", iterations, pixel, z, c, tooMany, tooFar, tileDone)
        }

        for count = 0; count < maxCount; count++ {
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
        //fmt.Printf("%05d ---- Counted #%v max: %v, Iterated #%v max: %v, pix: %v,  d: %v  z: %v c: %v, [[tooMany: %v, tooFar: %v, done: %v]]\n", totalLoops, count, maxCount, iterations, maxIterations, pixel, distSq, z, c, tooMany, tooFar, tileDone)

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
          color := getColor(iterations, maxIterations)
          //draw.Draw(tile.Img, tile.Img.Bounds(), &image.Uniform{color}, image.Point{}, draw.Src)
          //draw.Draw(tile.Img, image.Rectangle{image.Point{0, 0}, pixel}, &image.Uniform{color}, image.Point{}, draw.Src)

          onePixRect := image.Rect(pixel.X, pixel.Y, pixel.X + 1, pixel.Y + 1)
          draw.Draw(tile.Img, onePixRect, &image.Uniform{C: color}, image.Point{}, draw.Src)

          //fmt.Printf("Final #%v pixel: %v  Z: %v, C: %v, [[tooMany: %v, tooFar: %v, done: %v]]\n", iterations, color, z, c, tooMany, tooFar, tileDone)
          joe.dataChannels.messages <- fmt.Sprintf("%v: %v: %v %v", iterations, pixel, z, color)
        }

      }
    }

    if tileDone {
      fmt.Printf("Loops: [%v], Pixels: %v, fast: %v, fail: %v\n", totalLoops, totalPixels, totalFast, totalFails)
      tilechan <- tile
    }
  }()
  wg.Wait()

  return nil
}
