package corkscrew

/* Copyright © 2022 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

import (
  "fmt"
  "github.com/go-p5/p5"
  "image"
  "image/color"
  "math"
  "time"
)

type MandelOptions struct {
  Width       int
  Height      int

  // Either this one (part of set #1)
  PlotHeight  float32
  PlotWidth   float32
  PlotCenterX float32                  // defaults to origin if PlotWidth/PlotHeight are used
  PlotCenterY float32

  // Or this one (part of set #1)
  Left        float32
  Right       float32
  Top         float32
  Bottom      float32
}

var (
  count     = 0
  startTime = time.Now()
)

var opts              *MandelOptions
//var userDispWidth, userDispHeight int
//var userDomainWidth, userRangeHeight float32
var field           *Field
var joe             *Joe
var mandel          *MandelbrotTile
var tilechan         chan *Tile
var quit             chan struct{}

func init() {

  //quit = make(chan struct{})
  //
  //// Get from cli
  //userDispWidth = 800
  //userDispHeight = 600
  ////userDomainWidth = 100.0
  ////userRangeHeight = 100.0
  //userDomainWidth = 3.0
  //userRangeHeight = 2.0
  //
  //userRect := image.Rectangle{Max: image.Point{X: userDispWidth, Y: userDispHeight}}
  //field   = NewField(userRect, userDomainWidth, userRangeHeight)
  //joe     = NewJoe(field)
  //
  //fmin, fmax := field.FBounds(userDomainWidth, userRangeHeight)
  //
  //mandel  = NewMandelbrotTile(fmin, fmax, userRect, joe)
}

func ShowMandelbrotSet(opts_ MandelOptions) error {

  opts = GetMandelOpts(opts_)
  quit = make(chan struct{})

  // Get from cli
  //userDispWidth = 800
  //userDispHeight = 600
  //userDomainWidth = 100.0
  //userRangeHeight = 100.0
  //userDomainWidth = 3.0
  //userRangeHeight = 2.0

  userRect := image.Rectangle{Max: image.Point{X: opts.Width, Y: opts.Height}}
  field   = NewField(userRect, opts.Left, opts.Top, opts.Right, opts.Bottom)
  joe     = NewJoe(field)

  fmin, fmax := field.FBounds()

  mandel  = NewMandelbrotTile(fmin, fmax, userRect, joe)

  tilechan, err := joe.Run(quit)
  if err != nil {
    return err
  }

  err = mandel.Run(quit, field, tilechan)
  if err != nil {
    return err
  }

  p5.Run(setupP5, drawP5)
  return nil
}

func setupP5() {
  p5.Canvas(opts.Width, opts.Height)
  p5.Background(color.Gray{Y: 220})
}

func drawP5() {
  count++
  t := time.Now()
  elapsed := t.Sub(startTime).Seconds()
  sec := t.Second()

  joe.Render()

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
}


func GetMandelOpts(userOpts MandelOptions) *MandelOptions {

  opts := MandelOptions {
    //Width:          oneOrTheOther(userOpts.Width,   1200),
    //Height:         oneOrTheOther(userOpts.Height, 900),
    Width:          oneOrTheOther(userOpts.Width,   800),
    Height:         oneOrTheOther(userOpts.Height, 600),

    PlotWidth:      oneOrTheOtherF(userOpts.PlotWidth, 4.1),
    PlotHeight:     oneOrTheOtherF(userOpts.PlotHeight, 4.0),
    PlotCenterX:    oneOrTheOtherF(userOpts.PlotCenterX, 0.0),
    PlotCenterY:    oneOrTheOtherF(userOpts.PlotCenterY, 0.0),

    Left:           oneOrTheOtherF(userOpts.Left,  -2.1),
    Right:          oneOrTheOtherF(userOpts.Right, 2.0),
    Top:            oneOrTheOtherF(userOpts.Top, 2.0),
    Bottom:         oneOrTheOtherF(userOpts.Bottom, -2.0),
  }

  if opts.Left == 0.0 && opts.Right == 0.0 && opts.Top == 0.0 && opts.Bottom == 0.0 {
    halfWidth, halfHeight := opts.PlotWidth / 2.0, opts.PlotHeight / 2.0

    opts.Left     = opts.PlotCenterX - halfWidth
    opts.Right    = opts.PlotCenterX + halfWidth
    opts.Top      = opts.PlotCenterY + halfHeight
    opts.Bottom   = opts.PlotCenterY - halfHeight
  }

  return &opts
}

func (opts *MandelOptions) GetLeft() float32 {
 if opts.Left != 0 {
   return opts.Left
 } else if opts.PlotCenterX != 0 && opts.PlotWidth != 0 {
   return opts.PlotCenterX - opts.PlotWidth
 }

 return -2.1
}

func oneOrTheOther(one, other int) int {
  if one != 0 {
    return one
  }

  return other
}

func oneOrTheOtherF(one, other float32) float32 {
  if one != 0 {
    return one
  }

  return other
}
