package corkscrew

/* Copyright © 2022 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

// -------------------------------------------------------------------------------------------------------------------

func init() {
  Config.ShowGridLines   = true
  Config.ShowAxis        = true
  Config.ShowVertAxis    = true
  Config.ShowHorizAxis   = true
  Config.GridStyle       = byTens
}

// -------------------------------------------------------------------------------------------------------------------

type GridStyle int
const (
  byTens      GridStyle = iota
  byHalves
)
func (gs GridStyle) String() string {
  return [...]string{"by tens", "by halves"}[gs]
}

// -------------------------------------------------------------------------------------------------------------------

type ConfigOptions struct {
  ShowGridLines   bool
  ShowAxis        bool
  ShowVertAxis    bool
  ShowHorizAxis   bool
  GridStyle       GridStyle

  // TODO (axis/grid):
  // dark/light/transparent
  // thin stroke
  // labeled (units)
}
var Config ConfigOptions

// -------------------------------------------------------------------------------------------------------------------

type RuntimeOptions struct {
  multiThreaded       bool
  maxSplits           int
}
var runtimeOptions RuntimeOptions

func init() {
  runtimeOptions.multiThreaded = true
  runtimeOptions.maxSplits = 1
}

// -------------------------------------------------------------------------------------------------------------------

type DebugOptions struct {

}
var debugOptions DebugOptions

