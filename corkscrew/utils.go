package corkscrew

/* Copyright © 2022 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

import (
  "fmt"
  colorful "github.com/lucasb-eyer/go-colorful"
  "strings"
)

// -------------------------------------------------------------------------------------------------------------------

var palette []colorful.Color
var black, white colorful.Color

// -------------------------------------------------------------------------------------------------------------------

func init() {
  palette = gradientGen()

  black, _ = colorful.Hex("#000000")
  white, _ = colorful.Hex("#000000")
}

// -------------------------------------------------------------------------------------------------------------------

func appendTo(existing, msg string) string {
  if len(existing) > 0 {
    return existing + ", " + msg
  }
  return msg
}

// -------------------------------------------------------------------------------------------------------------------

func appendsTo(existing, sep, msg string) string {
  if len(existing) > 0 {
    if len(msg) > 0 {
      return existing + sep + msg
    }
    return existing
  }
  return msg
}

// -------------------------------------------------------------------------------------------------------------------

func getColor(iterations, maxIterations int) colorful.Color {
  percentage := float32(iterations) / float32(maxIterations)
  index := int(percentage * float32(len(palette)))

  var result colorful.Color

  if index >= len(palette) {
    result = black
  } else {
    result = palette[index]
  }

  //fmt.Printf("[%06d] %v idx: %v Picked color: %v\n", iterations, percentage, index, result)
  return result
}

// -------------------------------------------------------------------------------------------------------------------

func asserter(test bool) bool {
  if !test {
    breakout("", true)
  }
  return !test
}

// -------------------------------------------------------------------------------------------------------------------

func assertMsg(test bool, msg string) {
  if !test {
    breakout(msg, false)
  }
}

// -------------------------------------------------------------------------------------------------------------------

func assert(test bool) {
  if !test {
    breakout("", false)
  }
}

// -------------------------------------------------------------------------------------------------------------------

func breakout(msg string, quiet bool) {
  if !quiet {
    fmt.Printf("  ------------ BREAKOUT!! %v !!\n", msg)
  }
}

// -------------------------------------------------------------------------------------------------------------------

func BuildStringMap(ss []string) map[string]struct{} {
  m := map[string]struct{}{}

  for _, s := range ss {
    m[s] = struct{}{}
  }

  return m
}

// -------------------------------------------------------------------------------------------------------------------

func BuildStringMapStr(s string) map[string]struct{} {
  ss := strings.Split(s, ",")
  return BuildStringMap(ss)
}

// -------------------------------------------------------------------------------------------------------------------

func intersection(a, b map[string]struct{}) map[string]struct{} {
  in := map[string]struct{}{}
  for s, _ := range a {
    if _, ok := b[s]; ok {
      in[s] = struct{}{}
    }
  }
  return in
}

// -------------------------------------------------------------------------------------------------------------------

func union(a, b map[string]struct{}) map[string]struct{} {
  un := map[string]struct{}{}
  for s, _ := range a {
    un[s] = struct{}{}
  }
  for s, _ := range b {
    un[s] = struct{}{}
  }
  return un
}

// -------------------------------------------------------------------------------------------------------------------

func hasIntersection(a, b map[string]struct{}) bool {
  in := intersection(a, b)
  return len(in) > 0
}




