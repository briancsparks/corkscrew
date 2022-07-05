package cmd

/* Copyright © 2022 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

import (
	"fmt"
  "github.com/briancsparks/corkscrew/corkscrew"

  "github.com/spf13/cobra"
)

var opts corkscrew.MandelOptions

// mandelCmd represents the mandel command
var mandelCmd = &cobra.Command{
	Use:   "mandel",
	Short: "Mandelbrot set",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("mandel called")
	},
}

func init() {
	rootCmd.AddCommand(mandelCmd)

  width := mandelCmd.Flags().IntP("width", "w", 1200, "width")
  height := mandelCmd.Flags().Int("height", 800, "height")

  plotWidth  := mandelCmd.Flags().Float32("PlotWidth", 4.1, "The width of the data to plot")
  plotHeight := mandelCmd.Flags().Float32("PlotHeight", 4.0, "The height of the data to plot")
  centerx := mandelCmd.Flags().Float32("centerx", 0, "Put the center")
  centery := mandelCmd.Flags().Float32("centery", 0, "Put the center")

  left := mandelCmd.Flags().Float32P("left", "l", -2.1, "The left coordinate")
  right := mandelCmd.Flags().Float32P("right", "r", 1.2, "The right coordinate")
  top := mandelCmd.Flags().Float32P("top", "t", 1.2, "The top coordinate")
  bottom := mandelCmd.Flags().Float32P("bottiom", "b", -1.2, "The bottom coordinate")

  corkscrew.ShowMandelbrotSet(corkscrew.MandelOptions{
    Width:        *width,
    Height:       *height,

    // Either this one (part of set #1)
    PlotWidth:    *plotWidth,
    PlotHeight:   *plotHeight,
    PlotCenterX:  *centerx,
    PlotCenterY:  *centery,

    // Or this one (part of set #1)
    Left:         *left,
    Right:        *right,
    Top:          *top,
    Bottom:       *bottom,
  })

  // TODO:
  // * Axes, scrollbars, etc.

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// mandelCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// mandelCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
