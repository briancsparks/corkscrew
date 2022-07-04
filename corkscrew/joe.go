package corkscrew

/* Copyright © 2022 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

import (
  "github.com/go-p5/p5"
  "sync"
)

type Joe struct {
  Field *Field
  Tiles []*Tile

  lock sync.RWMutex
}

func NewJoe(f *Field) *Joe {
  j := &Joe{
    Field: f,
  }

  j.Tiles = append(j.Tiles, NewTile(70, 70, 5.0, 5.0, 0.0, 0.0, f))

  return j
}

func (j *Joe) Run(quit chan struct{}) (chan *Tile, error) {
  ch := make(chan *Tile)

  go func() {
    select {
    case t := <-ch:
      //j.Tiles = append(j.Tiles, t)
      j.AppendTile(t)

    case <- quit:
      break
    }
  }()

  return ch, nil
}

func (j *Joe) AppendTile(t *Tile) {
  j.lock.Lock()
  defer j.lock.Unlock()

  j.Tiles = append(j.Tiles, t)
}

func (j *Joe) Render() {
  // Get the lock for the operation
  j.lock.RLock()
  defer j.lock.RUnlock()

  for _, tile := range j.Tiles {
    p5.DrawImage(tile.Img, float64(tile.Min.X), float64(tile.Min.Y))
  }
}

