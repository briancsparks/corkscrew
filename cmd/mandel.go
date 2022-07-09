package cmd

/* Copyright © 2022 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

import (
	"fmt"
  "github.com/briancsparks/corkscrew/corkscrew"

  "github.com/spf13/cobra"
)

var width, height int
var plotWidth, plotHeight, plotRadius, centerx, centery float32
var left, right, top, bottom float32

// mandelCmd represents the mandel command
var mandelCmd = &cobra.Command{
	Use:   "mandel",
	Short: "Mandelbrot set",

	Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("mandel called")

    _ = corkscrew.ShowMandelbrotSet(corkscrew.MandelOptions{
      Width:        width,
      Height:       height,

      // Either this one (part of set #1)
      //PlotWidth:    plotWidth,
      //PlotHeight:   plotHeight,
      PlotRadius:   plotRadius,
      PlotCenterX:  centerx,
      PlotCenterY:  centery,

      // Or this one (part of set #1)
      Left:         left,
      Right:        right,
      Top:          top,
      Bottom:       bottom,
    })
	},
}

func init() {
	rootCmd.AddCommand(mandelCmd)

  mandelCmd.Flags().IntVar(&width, "width", 1200, "width")
  mandelCmd.Flags().IntVar(&height, "height", 800, "height")

  //mandelCmd.Flags().Float32Var(&plotWidth, "PlotWidth", 4.1, "The width of the data to plot")
  //mandelCmd.Flags().Float32Var(&plotHeight, "PlotHeight", 4.0, "The height of the data to plot")
  mandelCmd.Flags().Float32VarP(&plotRadius, "radius", "r", 0, "Put the radius")
  mandelCmd.Flags().Float32VarP(&centerx, "centerx", "x", 0, "Put the center")
  mandelCmd.Flags().Float32VarP(&centery, "centery", "y", 0, "Put the center")


  mandelCmd.Flags().Float32Var(&left, "left", -2.1, "The left coordinate")
  mandelCmd.Flags().Float32Var(&right, "right", 1.2, "The right coordinate")
  mandelCmd.Flags().Float32Var(&top, "top", 1.2, "The top coordinate")
  mandelCmd.Flags().Float32Var(&bottom, "bottom", -1.2, "The bottom coordinate")

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
