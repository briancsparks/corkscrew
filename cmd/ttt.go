package cmd

/* Copyright © 2022 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

import (
  "fmt"

  "github.com/briancsparks/corkscrew/corkscrew"
  "github.com/spf13/cobra"
)

// tttCmd represents the ttt command
var tttCmd = &cobra.Command{
  Use:   "ttt",
  Short: "Run the currently-in-development command",

  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("ttt called")
    _ = corkscrew.ShowMain()
  },
}

func init() {
  rootCmd.AddCommand(tttCmd)

  // Here you will define your flags and configuration settings.

  // Cobra supports Persistent Flags which will work for this command
  // and all subcommands, e.g.:
  // tttCmd.PersistentFlags().String("foo", "", "A help for foo")

  // Cobra supports local flags which will only run when this command
  // is called directly, e.g.:
  // tttCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
