package corkscrew

/* Copyright © 2022 sparksb -- MIT (see LICENSE file) */
/* Copyright © 2022 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

import (
	"fmt"
)

type DebugIt struct {
  ID            int32
  enabled       bool
  assoc         string

  enabledTags   map[string]struct{}

  info          map[string]string
}

// --------------------------------------------------------------------------------------------------------------------

func NewDebugIt(ID int32, assoc string) *DebugIt {
  d := &DebugIt{ID: ID, assoc: assoc}
  d.enabled = true
  d.info  = map[string]string{}
  d.enabledTags = map[string]struct{}{}

  //d.enabledTags["begin"] = struct{}{}
  //d.enabledTags["allPoints"] = struct{}{}
  //d.enabledTags["end"] = struct{}{}

  return d
}

// --------------------------------------------------------------------------------------------------------------------

func (dbg *DebugIt) isEnabled(tags *map[string]struct{}) bool {
  if tags == nil {
    return dbg.enabled
  }
  return hasIntersection(*tags, dbg.enabledTags)
}

// --------------------------------------------------------------------------------------------------------------------

func (dbg *DebugIt) Enable() {
  dbg.enabled = true
}

// --------------------------------------------------------------------------------------------------------------------

func (dbg *DebugIt) Disable() {
  dbg.enabled = false
}

// --------------------------------------------------------------------------------------------------------------------

func (dbg *DebugIt) Printf(format string, a ...interface{}) {
  if dbg.enabled {
    dbg.fmtPrintf(fmt.Sprintf("[%2d]: %s", dbg.ID, format), a...)
  }
}

// --------------------------------------------------------------------------------------------------------------------

func (dbg *DebugIt) Printfln(format string, a ...interface{}) {
  if dbg.enabled {
    dbg.fmtPrintf(fmt.Sprintf("[%2d]: %s\n", dbg.ID, format), a...)
  }
}

// --------------------------------------------------------------------------------------------------------------------

func (dbg *DebugIt) fmtPrintf(format string, a ...interface{}) {
  fmt.Printf(fmt.Sprintf("[%2d]: %s", dbg.ID, format), a...)
}

// --------------------------------------------------------------------------------------------------------------------

func (dbg *DebugIt) fmtPrintfln(format string, a ...interface{}) {
  fmt.Printf(fmt.Sprintf("[%2d]: %s\n", dbg.ID, format), a...)
}

// --------------------------------------------------------------------------------------------------------------------

func PrintfTagged(dbg *DebugIt, tags map[string]struct{}, format string, a ...interface{}) {
  if dbg.isEnabled(&tags) {
    dbg.fmtPrintf(format, a...)
  }
}

// --------------------------------------------------------------------------------------------------------------------

func PrintflnTagged(dbg *DebugIt, tags map[string]struct{}, format string, a ...interface{}) {
  if dbg.isEnabled(&tags) {
    dbg.fmtPrintfln(format, a...)
  }
}

// --------------------------------------------------------------------------------------------------------------------

func (dbg *DebugIt) BuildTaggedPrintf(tags string) func (format string, a ...interface{}) {
  m := BuildStringMapStr(tags)
  return func(format string, a ...interface{}) {
    PrintfTagged(dbg, m, format, a...)
  }
}

// --------------------------------------------------------------------------------------------------------------------

func (dbg *DebugIt) BuildTaggedPrintfln(tags string) func (format string, a ...interface{}) {
  m := BuildStringMapStr(tags)
  return func(format string, a ...interface{}) {
    PrintflnTagged(dbg, m, format, a...)
  }
}

// --------------------------------------------------------------------------------------------------------------------

func (dbg *DebugIt) PrintfReport(format string, a ...interface{}) {
  fmt.Printf(fmt.Sprintf("[%2d]: %s", dbg.ID, format), a...)
}

// --------------------------------------------------------------------------------------------------------------------

func (dbg *DebugIt) PrintflnReport(format string, a ...interface{}) {
  fmt.Printf(fmt.Sprintf("[%2d]: %s\n", dbg.ID, format), a...)
}

// --------------------------------------------------------------------------------------------------------------------



