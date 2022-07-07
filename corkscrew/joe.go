package corkscrew

/* Copyright © 2022 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

import (
  "github.com/go-p5/p5"
  "sync"
)

type Joe struct {
  Field *Field
  Tiles []*Tile
  MainTile *Tile
  message string

  dataChannels DataChannels

  lock sync.RWMutex
}

type DataChannels struct {
  tiles     chan *Tile
  messages  chan string
}

func NewJoe(f *Field) *Joe {
  j := &Joe{
    Field: f,
  }
  j.dataChannels = DataChannels{
    tiles: make(chan *Tile),
    messages: make(chan string),
  }

  j.Tiles = append(j.Tiles, NewTile(70, 70, 5.0, 5.0, 0.0, 0.0, f))

  return j
}

func (j *Joe) Run(quit chan struct{}) (chan *Tile, error) {
  ch := make(chan *Tile)


  go func() {
    for {
      select {
      case t := <-ch:
        //j.Tiles = append(j.Tiles, t)
        //j.AppendTile(t)
        j.SetMainTile(t)

      case msg := <-j.dataChannels.messages:
        j.message = msg

      case <-quit:
        break
      }
    }
  }()

  return ch, nil
}

func (j *Joe) SetMainTile(t *Tile) {
  j.lock.Lock()
  defer j.lock.Unlock()

  j.MainTile = t
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
  if j.MainTile != nil {
    p5.DrawImage(j.MainTile.Img, float64(j.MainTile.Min.X), float64(j.MainTile.Min.Y))
  }

  //p5.TextSize(24)
  //p5.Text(fmt.Sprintf("%s", j.message), 10, 300)
}

