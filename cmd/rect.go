package cmd

/* Copyright © 2022 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

import (
	"fmt"
  "github.com/briancsparks/corkscrew/corkscrew"

  "github.com/spf13/cobra"
)

var left, right, top, bottom float32

// rectCmd represents the rect command
var rectCmd = &cobra.Command{
	Use:   "rect",
	Short: "Show the Mandelbrot set - area given by rectangle",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("rect called")

    _ = corkscrew.ShowMandelbrotSet(corkscrew.MandelOptions{
      Width:        width,
      Height:       height,

      Left:         left,
      Right:        right,
      Top:          top,
      Bottom:       bottom,
    })
	},
}

func init() {
	mandelCmd.AddCommand(rectCmd)

  rectCmd.Flags().Float32Var(&left, "left", -2.1, "The left coordinate")
  rectCmd.Flags().Float32Var(&right, "right", 1.2, "The right coordinate")
  rectCmd.Flags().Float32Var(&top, "top", 1.2, "The top coordinate")
  rectCmd.Flags().Float32Var(&bottom, "bottom", -1.2, "The bottom coordinate")
}
