package cmd

/* Copyright © 2022 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

import (
	"fmt"
  "github.com/briancsparks/corkscrew/corkscrew"

  "github.com/spf13/cobra"
)

var width, height int
var plotRadius, centerx, centery float32
var where string

// mandelCmd represents the mandel command
var mandelCmd = &cobra.Command{
	Use:   "mandel",
	Short: "Mandelbrot set",

	Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("mandel called")

    if where != "center" {
      switch where {
      case "a":
        centerx = -0.7463
        centery = 0.1102
        plotRadius = 0.005
      case "b":
        centerx = -0.16
        centery = 1.0405
        plotRadius = 0.026
      }
    }

    _ = corkscrew.ShowMandelbrotSet(corkscrew.MandelOptions{
      Width:        width,
      Height:       height,

      PlotRadius:   plotRadius,
      PlotCenterX:  centerx,
      PlotCenterY:  centery,
    })
	},
}

func init() {
	rootCmd.AddCommand(mandelCmd)

  mandelCmd.PersistentFlags().IntVar(&width, "width", 1200, "width")
  mandelCmd.PersistentFlags().IntVar(&height, "height", 800, "height")

  mandelCmd.Flags().Float32VarP(&plotRadius, "radius", "r", 2.8, "Put the radius")
  mandelCmd.Flags().Float32VarP(&centerx, "centerx", "x", 0.0, "Put the center")
  mandelCmd.Flags().Float32VarP(&centery, "centery", "y", 0.0, "Put the center")

  mandelCmd.Flags().StringVar(&where, "where", "center", "Where to start plot [a-e]")

  //X = -0.7463
  //Y = 0.1102
  //R = 0.005
  //
  //X = -0.7453
  //Y = 0.1127
  //R = 6.5E-4
  //
  //X = -0.74529
  //Y = 0.113075
  //R = 1.5E-4
  //
  //X = -0.745428
  //Y = 0.113009
  //R = 3.0E-5
  //
  //X = -0.16
  //Y = 1.0405
  //R = 0.026


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
