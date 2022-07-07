package corkscrew

/* Copyright © 2022 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

import (
  "github.com/go-p5/p5"
  "image"
  "image/color"
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

var mandelOpts      MandelOptions
var field           *Field
var joe             *Joe
var mandel          *MandelbrotTile
var tilechan         chan *Tile
var quit             chan struct{}

func init() {
}

func ShowMandelbrotSet(opts_ MandelOptions) error {

  mandelOpts = GetMandelOpts(opts_)
  quit = make(chan struct{})

  userRect := image.Rectangle{Max: image.Point{X: mandelOpts.Width, Y: mandelOpts.Height}}
  field   = NewField(userRect, mandelOpts.Left, mandelOpts.Top, mandelOpts.Right, mandelOpts.Bottom)
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
  p5.Canvas(mandelOpts.Width, mandelOpts.Height)
  p5.Background(color.Gray{Y: 220})
}

// ------------------------------------------------------------------------------------------------------------------

func drawP5() {
  count++

  //t := time.Now()
  //elapsed := t.Sub(startTime).Seconds()
  //sec := t.Second()

  field.Render()
  joe.Render()
  field.RenderLast()

}

// ------------------------------------------------------------------------------------------------------------------

func GetMandelOpts(userOpts MandelOptions) MandelOptions {

  opts := MandelOptions {
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

  //func NormalizeRealm(realmRect Rec2, rec image.Rectangle) Rec2
  realmRect := R2(opts.Left, opts.Top, opts.Right, opts.Bottom)
  fixed     := NormalizeRealm(realmRect, image.Rect(0, 0, opts.Width, opts.Height))
  opts.Left, opts.Top, opts.Right, opts.Bottom = fixed.Parts()
  
  return opts
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
