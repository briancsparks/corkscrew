package corkscrew

/* Copyright © 2022 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

import (
  "github.com/go-p5/p5"
  "image"
  "sync"
)

// --------------------------------------------------------------------------------------------------------------------

type Joe struct {
  Field *Field
  Tiles []*Tile
  MainTile *Tile
  IdTiles   map[int32]*Tile
  message string

  bg          *image.RGBA

  dataChannels DataChannels

  lock sync.RWMutex
}

// --------------------------------------------------------------------------------------------------------------------

type DataChannels struct {
  tiles     chan *Tile
  messages  chan string
}

// --------------------------------------------------------------------------------------------------------------------

func NewJoe(f *Field) *Joe {
  j := &Joe{
    Field: f,
  }
  j.dataChannels = DataChannels{
    tiles: make(chan *Tile),
    messages: make(chan string),
  }
  j.IdTiles = map[int32]*Tile{}
  j.bg = image.NewRGBA(f.Bounds)

  //j.Tiles = append(j.Tiles, NewTile(70, 70, 5.0, 5.0, 0.0, 0.0, f))

  return j
}

// --------------------------------------------------------------------------------------------------------------------

func (j *Joe) Run(quit chan struct{}) (chan *Tile, error) {
  tilechan := make(chan *Tile)

  // TODO: use wg

  go func() {
    for {
      select {
      case tile := <-tilechan:
        //j.Tiles = append(j.Tiles, tile)
        //j.AppendTile(tile)
        //j.SetMainTile(tile)
        j.SetIdTile(tile)

      case msg := <-j.dataChannels.messages:
        j.message = msg

      case <-quit:
        break
      }
    }
  }()

  return tilechan, nil
}

// --------------------------------------------------------------------------------------------------------------------

func (j *Joe) SetIdTile(tile *Tile) {
  j.lock.Lock()
  defer j.lock.Unlock()

  //fmt.Printf("[%3d]: %v\n", tile.ID, tile.Rect)
  j.IdTiles[tile.ID] = tile
}

// --------------------------------------------------------------------------------------------------------------------

func (j *Joe) SetMainTile(t *Tile) {
  j.lock.Lock()
  defer j.lock.Unlock()

  j.MainTile = t
}

// --------------------------------------------------------------------------------------------------------------------

func (j *Joe) AppendTile(t *Tile) {
  j.lock.Lock()
  defer j.lock.Unlock()

  j.Tiles = append(j.Tiles, t)
}

// --------------------------------------------------------------------------------------------------------------------

func (j *Joe) RenderBg() {

  // Get the lock for the operation
  j.lock.RLock()
  defer j.lock.RUnlock()

  if len(j.IdTiles) > 0 {
    for _, tile := range j.IdTiles {
      //draw.Draw(j.bg, tile.Rect, tile.Img, image.Point{}, draw.Src)
      p5.DrawImage(tile.Img, float64(tile.Rect.Min.X), float64(tile.Rect.Min.Y))
    }
    //p5.DrawImage(j.bg, 0.0, 0.0)
  }

}

// --------------------------------------------------------------------------------------------------------------------

func (j *Joe) Render() {
  // Get the lock for the operation
  j.lock.RLock()
  defer j.lock.RUnlock()

  for _, tile := range j.Tiles {
    p5.DrawImage(tile.Img, float64(tile.Rect.Min.X), float64(tile.Rect.Min.Y))
  }
  if j.MainTile != nil {
    p5.DrawImage(j.MainTile.Img, float64(j.MainTile.Rect.Min.X), float64(j.MainTile.Rect.Min.Y))
  }

  //p5.TextSize(24)
  //p5.Text(fmt.Sprintf("%s", j.message), 10, 300)
}

