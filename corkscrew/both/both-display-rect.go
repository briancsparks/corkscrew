package both

/* Copyright © 2022 sparksb -- MIT (see LICENSE file) */
/* Copyright © 2022 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

type DisplayRect struct {
  Min, Max DisplayPt
}

// -----------------------------------------------------------------------------------------------------------

func Rect(left,top,right,bottom int) DisplayRect {
  left, top, right, bottom = NormalizeDisplayRect(left, top, right, bottom)
  return DisplayRect{Min: DisplayPt{left,top}, Max: DisplayPt{right, bottom}}
}

// -----------------------------------------------------------------------------------------------------------

func RectPts(min, max DisplayPt) DisplayRect {
  return Rect(min.X, min.Y, max.X, max.Y)
}

// -----------------------------------------------------------------------------------------------------------

func (r *DisplayRect) Dx() int {
  return r.Max.X - r.Min.X
}

// -----------------------------------------------------------------------------------------------------------

func (r *DisplayRect) Dy() int {
  return r.Max.Y - r.Min.Y
}

// -----------------------------------------------------------------------------------------------------------

func (r *DisplayRect) OffsetX(x int) int {
  return x - r.Min.X
}

// -----------------------------------------------------------------------------------------------------------

func (r *DisplayRect) OffsetY(y int) int {
  return y - r.Min.Y
}

// -----------------------------------------------------------------------------------------------------------

func (r *DisplayRect) PercentX(x int) float64 {
  return float64(r.OffsetX(x)) / float64(r.Dx())
}

// -----------------------------------------------------------------------------------------------------------

func (r *DisplayRect) PercentY(y int) float64 {
  return float64(r.OffsetY(y)) / float64(r.Dy())
}

// -----------------------------------------------------------------------------------------------------------

func (r *DisplayRect) GetPtPercent(percentX, percentY float64) DisplayPt {
  x := float64(r.Min.X) + percentX * float64(r.Dx())
  y := float64(r.Min.Y) + percentY * float64(r.Dy())

  return DisplayPt{int(x), int(y)}
}

// -----------------------------------------------------------------------------------------------------------

