package corkscrew

/* Copyright © 2022 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */


type ConfigOptions struct {
  ShowGridLines   bool
  ShowVertAxis    bool
  ShowHorizAxis   bool
}
var configOptions ConfigOptions

func init() {
  configOptions.ShowGridLines = true
  configOptions.ShowVertAxis = true
  configOptions.ShowHorizAxis = true
}

