package corkscrew

/* Copyright © 2022 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

import (
  "fmt"
  "image"
  "image/draw"
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

func (m *MandelbrotTile) Run(quit chan struct{}, tileMaker tileMaker, /*tilechan*/ _ chan *Tile) error {

  wg := sync.WaitGroup{}

  wg.Add(1)
  go func() {
    wg.Done()

    //z := 0+0i
    z := complex(float32(0.0), float32(0.0))
    c := complex(float32(0.0), float32(0.0))

    maxCount := 150
    iterations := 0
    maxIterations := 1000
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

    fetchNext := true

    for {
      select {
      case <-quit:
        break

      default:
        iterations++
        iterations--
        break
      }

      // If we need to fetch another to work on...
      if fetchNext {
        // Get initial
        v2, pixel = tile.sub.Next()
        c = complex(v2.X, v2.Y)
        iterations = 0
        fmt.Printf("--Putting #%v pixel: %v  z: %v c: %v\n", iterations, pixel, z, c)
      }

      var badDistance, tooMany bool
      var doneWithCurrent bool

      for count := 0; count < maxCount; count++ {
        z = z*z + c
        iterations++

        if iterations > maxIterations {
          tooMany = true
          break
        }

        // How far are we?
        re, im := real(z), imag(z)
        distSq = re*re+im*im
        if distSq > thresholdDist {
          // We are off in the weeds
          badDistance = true
          break
        }
      }
      fmt.Printf("----Counted #%v pixel: %v  d: %v\n", count, maxCount, distSq)

      if tooMany {
        // Its stuck
        doneWithCurrent = true
        fetchNext = true
      } else if badDistance {
        // Too far away, will never come back
        doneWithCurrent = true
        fetchNext = true
      } else {
        // Ran out of time for this run. Not an error, just loop around
      }

      // If we are done with this pixel, report it or whatever
      if doneWithCurrent {
        color := getColor(iterations)
        //draw.Draw(tile.Img, tile.Img.Bounds(), &image.Uniform{color}, image.Point{}, draw.Src)
        draw.Draw(tile.Img, image.Rectangle{image.Point{0,0}, pixel}, &image.Uniform{color}, image.Point{}, draw.Src)
        fmt.Printf("Putting #%v pixel: %v  Z: %v, C: %v\n", iterations, color, z, c)
        joe.dataChannels.messages <- fmt.Sprintf("%v: %v: %v %v", iterations, pixel, z, color)
      }

    }


  }()
  wg.Wait()

  return nil
}
