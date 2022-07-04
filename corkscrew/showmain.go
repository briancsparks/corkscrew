package corkscrew

/* Copyright © 2022 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

import (
  "fmt"
  "image"
  "image/color"
  "image/draw"
  "math"
  "time"

  "github.com/go-p5/p5"
  colorful "github.com/lucasb-eyer/go-colorful"
)

var (
  tile      = image.NewRGBA(image.Rect(0, 0, 100, 100))
  count     = 0
  startTime = time.Now()
)

func init() {
  c := colorful.WarmColor()
  draw.Draw(tile, tile.Bounds(), &image.Uniform{C: color.RGBA{R: uint8(c.R * 255), G: uint8(c.G * 255), B: uint8(c.B * 255), A: 255}}, image.ZP, draw.Src)
}

func ShowMain() {
  p5.Run(setupP5, drawP5)
}

func setupP5() {
  p5.Canvas(800, 800)
  p5.Background(color.Gray{Y: 220})
}

func drawP5() {
  count++
  t := time.Now()
  elapsed := t.Sub(startTime).Seconds()

  sec := t.Second()
  clockStart := -(math.Pi / 2)

  p5.StrokeWidth(2)
  p5.Fill(color.RGBA{R: 255, A: 208})
  p5.Ellipse(50, 50, 80, 80)

  p5.Fill(color.RGBA{B: 255, A: 208})
  p5.Quad(50, 50, 80, 50, 80, 120, 60, 120)

  p5.Fill(color.RGBA{G: 255, A: 208})
  p5.Rect(200, 200, 50, 100)

  p5.Fill(color.RGBA{G: 255, A: 208})
  p5.Triangle(100, 100, 120, 120, 80, 120)

  p5.TextSize(24)
  p5.Text(fmt.Sprintf("%d - %v", count, float64(count)/elapsed), 10, 300)

  p5.Stroke(color.Black)
  p5.StrokeWidth(5)
  p5.Arc(300, 100, 80, 80, clockStart, clockStart+(float64(sec)/60.0)*2.0*math.Pi)

  p5.DrawImage(tile, 20, 20)
}
