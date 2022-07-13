package corkscrew

/* Copyright © 2022 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

import (
	"fmt"
  "strconv"
  "strings"
)

// -------------------------------------------------------------------------------------------------------------------

func FindBestLineForLabel(min, max float64) []float64 {
  if min < 0.0 && max > 0.0 {
    var result []float64
    right := FindBestLineForLabel(0.0, max)
    left := FindBestLineForLabel(0.0, -min)
    for i := len(left)-1; i >= 0; i-- {
      if -left[i] != right[0] {
        result = append(result, -left[i])
      }
    }
    return append(result, right...)
  }

  smin := fmt.Sprintf("%040.22f", min)
  smax := fmt.Sprintf("%040.22f", max)
  commonDigits := smin[:]
  point := min + (max - min) / 2.0
  pointStr := fmt.Sprintf("%040.22f\n", point)
  _,_,_,_=smin, commonDigits,smax,pointStr

  var result []float64
  var indexDiff int
  for indexDiff = 0; indexDiff < len(smin); indexDiff++ {
    if smin[indexDiff] != smax[indexDiff] {
      break
    }
  }
  //index := indexDiff - 1

  commonDigits = commonDigits[:indexDiff]
  for i := 0; i < 10; i++ {
    //oneResult2 := commonDigits[:index] + string([]byte{byte(i + '0')})
    x, _ := strconv.ParseFloat(commonDigits + string([]byte{byte(i + '0')}), 64)
    if min <= x && x <= max {
      result = append(result, x)
    }
  }

  return result
}

// -------------------------------------------------------------------------------------------------------------------

func FindBestLineForLabelX(min, max float64) float64 {
  point := min + (max - min) / 2.0
  best := point
  whole := false

  for IsWithinRange(point, min, max) && !whole {
    best = point
    fmt.Printf("best: %v\n", best)
    digits := fmt.Sprintf("%f", point)
    digits = strings.TrimRight(digits, ".0")
    whole = !strings.ContainsAny(digits, ".")
    newDigits := digits[:len(digits) - 1]
    if len(newDigits) == 0 {
      break
    }
    point, _ = strconv.ParseFloat(newDigits, 64)
  }
  return best
}

func IsWithinRange(x, min, max float64) bool {
  return min <= x && x <= max
}

