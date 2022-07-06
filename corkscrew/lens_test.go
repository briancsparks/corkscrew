package corkscrew

import (
  "image"
  "reflect"
  "testing"
)

func TestNormalizeRealm(t *testing.T) {
  type args struct {
    realmRect Rec2
    rec       image.Rectangle
  }
  tests := []struct {
    name string
    args args
    want Rec2
  }{
    //{
    //  name: "Test Normalize realmRect",
    //  args: args{
    //    realmRect: Rec2{Min: Vec2{-2.1,1.2}, Max: Vec2{1.2,-1.2}},
    //    rec: image.Rect(0, 0, 1200, 800),
    //  },
    //  want: Rec2{Vec2{-2.25, 1.2}, Vec2{1.35, -1.2}},
    //},
    {
      name: "Test Normalize realmRect to Y",
      args: args{
        realmRect: Rec2{Min: Vec2{-300,400}, Max: Vec2{300,-400}},
        rec: image.Rect(0, 0, 1200, 800),
      },
      want: Rec2{Min: Vec2{-600,400}, Max: Vec2{600,-400}},
    },
    {
      name: "Test Normalize realmRect to Y",
      args: args{
        realmRect: Rec2{Min: Vec2{-600,200}, Max: Vec2{600,-200}},
        rec: image.Rect(0, 0, 1200, 800),
      },
      want: Rec2{Min: Vec2{-600,400}, Max: Vec2{600,-400}},
    },
  }

  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      //got := NormalizeRealm(tt.args.realmRect, tt.args.rec)
      //roundX, roundY, roundX2, roundY2 := got.Parts()
      //roundX, roundY, roundX2, roundY2 = roundEm(roundX, roundY, roundX2, roundY2, 1000.0)
      //rounded := R2(roundX, roundY, roundX2, roundY2)
      //if !reflect.DeepEqual(rounded, tt.want) {

      if got := NormalizeRealm(tt.args.realmRect, tt.args.rec); !reflect.DeepEqual(got, tt.want) {
        t.Errorf("NormalizeRealm() = %v, want %v", got, tt.want)
      }
    })
  }
}

func roundEm(a, b, c, d float32, mul float32) (float32, float32, float32, float32) {
  return roundIt(a, mul), roundIt(b, mul), roundIt(c, mul), roundIt(d, mul)
}

func roundIt(x float32, mul float32) float32 {
  xi000 := int64(x * mul + 0.5)
  return float32(xi000) / mul
}

func TestRealmToScreen(t *testing.T) {
  type args struct {
    pt        Vec2
    realmRect Rec2
    rec       image.Rectangle
  }
  tests := []struct {
    name string
    args args
    want image.Point
  }{
    {
      name: "Easy First Test",
      args: args{
        pt: Vec2{1,1},
        realmRect: Rec2{Min: Vec2{0,0}, Max: Vec2{100,100}},
        rec: image.Rect(0, 0, 100, 100),
      },
      want: image.Point{X: 1, Y: 199},

    }, {
     name: "Main First Mandelbrot zoom level",
     args: args{
       pt: Vec2{-1.0,1.0},
       realmRect: Rec2{Min: Vec2{-2.25,1.2}, Max: Vec2{1.35,-1.2}},
       rec: image.Rect(0, 0, 1200, 800),
     },
     want: image.Point{X: 416, Y: 66},
    },
  }
  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      if got := RealmToScreen(tt.args.pt, tt.args.realmRect, tt.args.rec); !reflect.DeepEqual(got, tt.want) {
        t.Errorf("RealmToScreen() = %v, want %v", got, tt.want)
      }
    })
  }
}
