package corkscrew

/* Copyright © 2022 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

import (
  "image"
  "sync"
)

type FieldSplitter struct {
  Min, Max Vec2
  Bounds image.Rectangle

  first, second *MandelbrotTile

  joe *Joe
}

func NewFieldSplitter(numSplits int, min, max Vec2/*, x, y float64*/, bounds image.Rectangle, joe *Joe) *FieldSplitter {
  Validate(min, max)
  m := &FieldSplitter{
    Min:    min,
    Max:    max,
    Bounds: bounds,
  }
  m.joe = joe

  baseId := numSplits * 1000
  if m.Bounds.Dx() > m.Bounds.Dy() {
    midpoint := m.Bounds.Dx() / 2
    leftMax := image.Point{X: midpoint, Y: m.Bounds.Max.Y}
    rightMin := image.Point{X: midpoint, Y: m.Bounds.Min.Y}
    leftMaxRealm := ScreenToRealm4Pts(leftMax,  min, max, bounds.Min, bounds.Max)
    rightMinRealm := ScreenToRealm4Pts(rightMin, min, max, bounds.Min, bounds.Max)

    if numSplits <= 1 {
      m.first = NewMandelbrotTile(baseId + 0, min, leftMaxRealm, image.Rectangle{Min: bounds.Min, Max: leftMax}, joe)
      m.second = NewMandelbrotTile(baseId + 1, rightMinRealm, max, image.Rectangle{Min: rightMin, Max: bounds.Max}, joe)
    } else {
      _ = NewFieldSplitter(numSplits - 1, min, leftMaxRealm, image.Rectangle{Min: bounds.Min, Max: leftMax}, joe)
      _ = NewFieldSplitter(numSplits - 1, rightMinRealm, max, image.Rectangle{Min: rightMin, Max: bounds.Max}, joe)
    }

  } else {
    midpoint := m.Bounds.Dy() / 2
    topMax := image.Point{X: m.Bounds.Max.X, Y: midpoint}
    bottomMin := image.Point{X: m.Bounds.Min.X, Y: midpoint}
    topMaxRealm := ScreenToRealm4Pts(topMax, min, max, bounds.Min, bounds.Max)
    bottomMinRealm := ScreenToRealm4Pts(bottomMin, min, max, bounds.Min, bounds.Max)

    if numSplits <= 1 {
      m.first = NewMandelbrotTile(baseId + 0, min, topMaxRealm, image.Rectangle{Min: bounds.Min, Max: topMax}, joe)
      m.second = NewMandelbrotTile(baseId + 1, bottomMinRealm, max, image.Rectangle{Min: bottomMin, Max: bounds.Max}, joe)
    } else {
      _ = NewFieldSplitter(numSplits - 1, min, topMaxRealm, image.Rectangle{Min: bounds.Min, Max: topMax}, joe)
      _ = NewFieldSplitter(numSplits - 1, bottomMinRealm, max, image.Rectangle{Min: bottomMin, Max: bounds.Max}, joe)
    }
  }

  return m
}

func (m *FieldSplitter) Run(quit chan struct{}, tileMaker tileMaker, tilechan chan *Tile) error {

  wg := sync.WaitGroup{}

  wg.Add(1)
  go func() {
    wg.Done()

    _ = m.first.Run(quit, tileMaker, tilechan)
    _ = m.second.Run(quit, tileMaker, tilechan)

  }()
  wg.Wait()

  return nil
}

