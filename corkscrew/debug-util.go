package corkscrew

/* Copyright © 2022 sparksb -- MIT (see LICENSE file) */
/* Copyright © 2022 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

import (
	"fmt"
)

func debugit() {
	fmt.Printf("\n")
}

type DebugIt struct {
  ID            int32
  enabled       bool

  info          map[string]string
}

func NewDebugIt(ID int32) *DebugIt {
  d := &DebugIt{ID: ID}
  d.enabled = true
  d.info  = map[string]string{}
  return d
}

func (dbg *DebugIt) Enable() {
  dbg.enabled = true
}

func (dbg *DebugIt) Disable() {
  dbg.enabled = false
}

func (dbg *DebugIt) Printf(format string, a ...interface{}) {
  if dbg.enabled {
    fmt.Printf(fmt.Sprintf("[%2d]: %s", dbg.ID, format), a...)
  }
}

func (dbg *DebugIt) Printfln(format string, a ...interface{}) {
  if dbg.enabled {
    fmt.Printf(fmt.Sprintf("[%2d]: %s\n", dbg.ID, format), a...)
  }
}


func (dbg *DebugIt) PrintfReport(format string, a ...interface{}) {
  fmt.Printf(fmt.Sprintf("[%2d]: %s", dbg.ID, format), a...)
}

func (dbg *DebugIt) PrintflnReport(format string, a ...interface{}) {
  fmt.Printf(fmt.Sprintf("[%2d]: %s\n", dbg.ID, format), a...)
}



