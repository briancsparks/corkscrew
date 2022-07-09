package corkscrew

/* Copyright © 2022 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

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
  ShowVertAxis    bool
  ShowHorizAxis   bool
  GridStyle       GridStyle

  // TODO (axis/grid):
  // dark/light/transparent
  // thin stroke
  // labeled (units)
}
var configOptions ConfigOptions

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

// -------------------------------------------------------------------------------------------------------------------

func init() {
  configOptions.ShowGridLines   = true
  configOptions.ShowVertAxis    = true
  configOptions.ShowHorizAxis   = true
  configOptions.GridStyle       = byTens
}

