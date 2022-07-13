package corkscrew

import (
  "reflect"
  "testing"
)


func TestFindBestLineForLabel(t *testing.T) {
  type args struct {
    min float64
    max float64
  }
  tests := []struct {
    name string
    args args
    want []float64
  }{
    {
      name: "Finds one, too, three xyz",
      args: args{
        -2.4,
        1.2,
      },
      want: []float64{-2.0, -1.0, 0.0, 1.0},
    },
    {
      name: "Finds one, too, three",
      args: args{
        -0.2,
        0.3,
      },
      want: []float64{-0.2, -0.1, 0.0, 0.1, 0.2, 0.3},
    },
    {
      name: "Finds multiple",
      args: args{
        22.5555555534516,
        22.5555555534587,
      },
      want: []float64{22.555555553452, 22.555555553453, 22.555555553454, 22.555555553455, 22.555555553456, 22.555555553457, 22.555555553458},
    },
    {
      name: "Finds one, too",
      args: args{
        22.5555555534576,
        22.5555555534587,
      },
      want: []float64{22.555555553458},
    },
    {
     name: "Finds one",
     args: args{
       0.6,
       1.6,
     },
     want: []float64{1.0},
    },
    {
     name: "Finds two",
     args: args{
       1.5,
       2.5,
     },
     want: []float64{2.0},
    },
    {
     name: "Finds longer one",
     args: args{
       0.63483738,
       0.634848738,
     },
     want: []float64{0.63484},
    },
  }
  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      if got := FindBestLineForLabel(tt.args.min, tt.args.max); !reflect.DeepEqual(got, tt.want) {
        t.Errorf("FindBestLineForLabel() = %v, want %v", got, tt.want)
      }
    })
  }
}
