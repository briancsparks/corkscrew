package both

/* Copyright © 2022 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

type WorkRect struct {
  Min, Max WorkPt
}

// -----------------------------------------------------------------------------------------------------------

func WkRect(left,top,right,bottom float64) WorkRect {
  left, top, right, bottom = NormalizeWorkRect(left, top, right, bottom)
  return WorkRect{Min: WorkPt{left,top}, Max: WorkPt{right, bottom}}
}

// -----------------------------------------------------------------------------------------------------------

func WkRectPts(min, max WorkPt) WorkRect {
  return WkRect(min.X, min.Y, max.X, max.Y)
}

// -----------------------------------------------------------------------------------------------------------

func (r *WorkRect) Dx() float64 {
  return r.Max.X - r.Min.X
}

// -----------------------------------------------------------------------------------------------------------

func (r *WorkRect) Dy() float64 {
  return r.Min.Y - r.Max.Y
}

// -----------------------------------------------------------------------------------------------------------

func (r *WorkRect) MidpointX() float64 {
  return r.Max.X - r.Min.X
}

// -----------------------------------------------------------------------------------------------------------

func (r *WorkRect) MidpointY() float64 {
  return r.Min.Y - r.Max.Y
}

// -----------------------------------------------------------------------------------------------------------

func (r *WorkRect) OffsetX(x float64) float64 {
  return x - r.Min.X
}

// -----------------------------------------------------------------------------------------------------------

func (r *WorkRect) OffsetY(y float64) float64 {
  return y - r.Max.Y
}

// -----------------------------------------------------------------------------------------------------------

func (r *WorkRect) PercentX(x float64) float64 {
  return r.OffsetX(x) / r.Dx()
}

// -----------------------------------------------------------------------------------------------------------

func (r *WorkRect) PercentY(y float64) float64 {
  return r.OffsetY(y) / r.Dy()
}

// -----------------------------------------------------------------------------------------------------------

func (r *WorkRect) GetPtPercent(percentX, percentY float64) WorkPt {
  x := r.Min.X + percentX * r.Dx()
  y := r.Max.Y + percentY * r.Dy()

  return WorkPt{x, y}
}
