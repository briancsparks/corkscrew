package corkscrew

/* Copyright © 2022 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

import (
  "fmt"
  "github.com/briancsparks/corkscrew/corkscrew/both"
  "github.com/go-p5/p5"
  "image"
  "image/color"
  "sync"
)

// -------------------------------------------------------------------------------------------------------------------

var (
  cmd    *MandelBothCmd
)

// -------------------------------------------------------------------------------------------------------------------

type MandelConfig struct {
  displayWidth    int
  displayHeight   int

  centerx         float64
  centery         float64
  radius          float64
  offsetx         float64
  offsety         float64

  left            float64
  top             float64
  right           float64
  bottom          float64

  displayRect     image.Rectangle

  mandelImg      *image.RGBA
  mandelRect      image.Rectangle
}

// -------------------------------------------------------------------------------------------------------------------

func NewMandelConfig(displayWidth int, displayHeight int, centerx, centery, radius float64, offsetx, offsety float64) *MandelConfig {
  m := &MandelConfig{
    displayWidth: displayWidth,
    displayHeight: displayHeight,
    centerx:        centerx,
    centery:        centery,
    radius:         radius,
    offsetx:        offsetx,
    offsety:        offsety,
  }

  return m
}

// -------------------------------------------------------------------------------------------------------------------

type MandelBothCmd struct {
  config       *MandelConfig
  grid         *both.Both
  lock          sync.RWMutex
  mandelImg    *image.RGBA
  mandelRect    image.Rectangle

  previousDrawWorkBounds *both.WorkRect
}

// -------------------------------------------------------------------------------------------------------------------

func MandelBothMain(paramsIn *MandelConfig) error {
  params := figureOutConfig(paramsIn)

  cmd = &MandelBothCmd{}
  cmd.config = params

  params.displayRect = image.Rect(0, 0, params.displayWidth, params.displayHeight)
  params.mandelImg = image.NewRGBA(params.displayRect)
  data := both.MakeGrid(0,0, 0, params.displayWidth, params.displayHeight, params.left, params.top, params.right, params.bottom, both.SeeAll)
  cmd.grid = data

  var tilechan chan *MandelDataMessage = make(chan *MandelDataMessage)
  _= RunMandel(quit, tilechan, params, data)

  cmd.updater(quit, tilechan)

  p5.Run(mandelSetupP5, mandelDrawP5)
  return nil
}

func (c *MandelBothCmd) updater(quit chan struct{}, tilechan chan *MandelDataMessage) {

  go func() {
    for {
      select{
      case msg := <- tilechan:
        c.updateMsg(msg)

      case <- quit:
        break
      }
    }
  }()
}

// -------------------------------------------------------------------------------------------------------------------

func (c *MandelBothCmd) updateMsg(msg *MandelDataMessage) {
  c.lock.Lock()
  defer c.lock.Unlock()

  c.mandelImg = msg.Img
  c.mandelRect = msg.Rect
}

// -------------------------------------------------------------------------------------------------------------------

func mandelSetupP5() {
  p5.Canvas(cmd.config.displayWidth, cmd.config.displayHeight)
  p5.Background(color.Gray{Y: 220})
}

// ------------------------------------------------------------------------------------------------------------------

func mandelDrawP5() {
  cmd.DrawP5()
}

// -------------------------------------------------------------------------------------------------------------------

func figureOutConfig(paramsIn *MandelConfig) *MandelConfig {
  params := &MandelConfig{
    displayWidth:  oneOrTheOther(paramsIn.displayWidth, 1200),
    displayHeight: oneOrTheOther(paramsIn.displayHeight, 900),
    centerx:       paramsIn.centerx,
    centery:       paramsIn.centery,
    radius:        paramsIn.radius,
    offsetx:       paramsIn.offsetx,
    offsety:       paramsIn.offsety,
    left:          paramsIn.left,
    top:           paramsIn.top,
    right:         paramsIn.right,
    bottom:        paramsIn.bottom,
  }

  // We need ltrb
  if params.left == 0.0 && params.top == 0.0 && params.right == 0.0 && params.bottom == 0.0 {
    if params.radius != 0.0 {
      params.centerx  = params.centerx - params.offsetx
      params.centery  = params.centery - params.offsety

      params.left     = params.centerx - params.radius
      params.top      = params.centery + params.radius
      params.right    = params.centerx + params.radius
      params.bottom   = params.centery - params.radius
    }
  }

  if params.left == 0.0 && params.top == 0.0 && params.right == 0.0 && params.bottom == 0.0 {
    return nil
  }

  return params
}

// -------------------------------------------------------------------------------------------------------------------

func (m *MandelConfig) String() string {
  return fmt.Sprintf("(%v, (%v, %v)) -> [(x: %v, y: %v, r: %v (dx: %v, dy: %v)); (%v, %v, %v, %v)]",
      m.displayRect,
      m.displayWidth, m.displayHeight,
      m.centerx, m.centery, m.radius, m.offsetx, m.offsety,
      m.left, m.top, m.radius, m.bottom,
  )
}
