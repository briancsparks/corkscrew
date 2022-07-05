package corkscrew

/* Copyright © 2022 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

//
//var (
//  //tile      = image.NewRGBA(image.Rect(0, 0, 100, 100))
//  count     = 0
//  startTime = time.Now()
//
//)
//
//var userDispWidth, userDispHeight int
//var userDomainWidth, userRangeHeight float32
//var field           *Field
//var joe             *Joe
//var mandel          *MandelbrotTile
//var tilechan         chan *Tile
//var quit             chan struct{}
//
//func init() {
//  //c := colorful.WarmColor()
//  //draw.Draw(tile, tile.Bounds(), &image.Uniform{C: color.RGBA{R: uint8(c.R * 255), G: uint8(c.G * 255), B: uint8(c.B * 255), A: 255}}, image.ZP, draw.Src)
//
//  quit = make(chan struct{})
//
//  // Get from cli
//  userDispWidth = 800
//  userDispHeight = 600
//  //userDomainWidth = 100.0
//  //userRangeHeight = 100.0
//  userDomainWidth = 3.0
//  userRangeHeight = 2.0
//
//  userRect := image.Rectangle{Max: image.Point{X: userDispWidth, Y: userDispHeight}}
//  field   = NewField(userRect, userDomainWidth, userRangeHeight)
//  joe     = NewJoe(field)
//
//  fmin, fmax := field.FBounds(userDomainWidth, userRangeHeight)
//
//  mandel  = NewMandelbrotTile(fmin, fmax, userRect, joe)
//}
//
//func ShowMain() error {
//
//  tilechan, err := joe.Run(quit)
//  if err != nil {
//    return err
//  }
//
//  mandel.Run(quit, field, tilechan)
//
//  p5.Run(setupP5, drawP5)
//  return nil
//}
//
//func setupP5() {
//  p5.Canvas(1200, 900)
//  p5.Background(color.Gray{Y: 220})
//}
//
//func drawP5() {
//  count++
//  t := time.Now()
//  elapsed := t.Sub(startTime).Seconds()
//  sec := t.Second()
//
//  joe.Render()
//
//  clockStart := -(math.Pi / 2)
//
//  p5.StrokeWidth(2)
//  p5.Fill(color.RGBA{R: 255, A: 208})
//  p5.Ellipse(50, 50, 80, 80)
//
//  p5.Fill(color.RGBA{B: 255, A: 208})
//  p5.Quad(50, 50, 80, 50, 80, 120, 60, 120)
//
//  p5.Fill(color.RGBA{G: 255, A: 208})
//  p5.Rect(200, 200, 50, 100)
//
//  p5.Fill(color.RGBA{G: 255, A: 208})
//  p5.Triangle(100, 100, 120, 120, 80, 120)
//
//  p5.TextSize(24)
//  p5.Text(fmt.Sprintf("%d - %v", count, float64(count)/elapsed), 10, 300)
//
//  p5.Stroke(color.Black)
//  p5.StrokeWidth(5)
//  p5.Arc(300, 100, 80, 80, clockStart, clockStart+(float64(sec)/60.0)*2.0*math.Pi)
//}
