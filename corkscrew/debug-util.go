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

  info          map[string]string
}

func NewDebugIt(ID int32) *DebugIt {
  d := &DebugIt{ID: ID}
  d.info  = map[string]string{}
  return d
}



