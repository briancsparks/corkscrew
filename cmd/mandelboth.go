package cmd

/* Copyright © 2022 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

import (
  "fmt"
  "github.com/briancsparks/corkscrew/corkscrew"

  "github.com/spf13/cobra"
)

var offsetx, offsety float64
var centerx64, centery64, plotRadius64 float64

// mandelbothCmd represents the mandelboth command
var mandelbothCmd = &cobra.Command{
	Use:   "mandelboth",
	Short: "Fast Mandelbrot",

	Run: func(cmd *cobra.Command, args []string) {
    centerx64, centery64, plotRadius64, offsetx, offsety := location(where, centerx64, centery64, plotRadius64, offsetx, offsety)
    params := corkscrew.NewMandelConfig(width, height, centerx64, centery64, plotRadius64, offsetx, offsety)

    fmt.Printf("user params: %v\n", params)
    _ = corkscrew.MandelBothMain(params)
	},
}

func init() {
	rootCmd.AddCommand(mandelbothCmd)

  mandelbothCmd.PersistentFlags().IntVar(&width, "width", 1200, "width")
  mandelbothCmd.PersistentFlags().IntVar(&height, "height", 800, "height")

  mandelbothCmd.Flags().Float64VarP(&plotRadius64, "radius", "r", 1.2, "Put the radius")
  mandelbothCmd.Flags().Float64VarP(&centerx64, "centerx", "x", 0.0, "Put the center")
  mandelbothCmd.Flags().Float64VarP(&centery64, "centery", "y", 0.0, "Put the center")
  mandelbothCmd.Flags().Float64Var(&offsetx, "offsetx", 0, "Put the center")
  mandelbothCmd.Flags().Float64Var(&offsety, "offsety", 0, "Put the center")

  mandelbothCmd.Flags().StringVar(&where, "where", "center", "Where to start plot [a-e]")

}



func location(where string, centerx, centery, plotRadius float64, offsetx, offsety float64) (float64, float64, float64, float64, float64) {
  //centerx, centery, plotRadius:= 0.0, 0.0, 2.8

  if where != "center" {
    switch where {
    case "origin":
      centerx = 0.0
      centery = 0.0
      plotRadius = 1.2
      offsetx = 0.6
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

  return centerx, centery, plotRadius, offsetx, offsety
}
