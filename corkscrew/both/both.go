package both

/* Copyright © 2022 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

type Strategy int
const (
  SeeAll  Strategy = iota
  ZoomedInCropped
)
func (s Strategy) String() string {
  return [...]string{"SeeAll", "ZoomedInCropped"}[s]
}

// -----------------------------------------------------------------------------------------------------------------

type Both struct {
  Main  BothRect        /* The rect for the whole display */
}

// -----------------------------------------------------------------------------------------------------------------

type BothPt struct {
  Display DisplayPt
  Work    WorkPt
}

// -----------------------------------------------------------------------------------------------------------------

type BothRect struct {
  Display  DisplayRect
  Work     WorkRect
}

// -----------------------------------------------------------------------------------------------------------------

func (b *Both) WorkPtFor(pt DisplayPt) WorkPt {
  percentX := b.Main.Display.PercentX(pt.X)
  percentY := b.Main.Display.PercentY(pt.Y)
  return b.Main.Work.GetPtPercent(percentX, percentY)
}

// -----------------------------------------------------------------------------------------------------------------

func (b *Both) DisplayPtFor(pt WorkPt) DisplayPt {
  percentX := b.Main.Work.PercentX(pt.X)
  percentY := b.Main.Work.PercentY(pt.Y)
  return b.Main.Display.GetPtPercent(percentX, percentY)
}

// -----------------------------------------------------------------------------------------------------------------

type DisplayIterator struct {
  rects   *BothRect
  pt       BothPt

  dx, dy    float64
}

// -----------------------------------------------------------------------------------------------------------------

func (b *Both) GetDisplayIterator() DisplayIterator {
  di := DisplayIterator{
    rects: &b.Main,
    pt:     BothPt{
      Display:  b.Main.Display.Min,
      Work:     b.Main.Work.Min,
    },
  }

  unitsPerPixelX := b.Main.Work.Dx() / float64(b.Main.Display.Dx())
  unitsPerPixelY := b.Main.Work.Dy() / float64(b.Main.Display.Dy())
  di.dx, di.dy = unitsPerPixelX, unitsPerPixelY

  return di
}

// -----------------------------------------------------------------------------------------------------------------

func (di *DisplayIterator) Curr() *BothPt {
  return &di.pt
}

// -----------------------------------------------------------------------------------------------------------------

func (di *DisplayIterator) Next() *BothPt {
  di.pt.Display.X += 1
  if di.pt.Display.X >= di.rects.Display.Max.X {
    di.pt.Display.Y += 1
    if di.pt.Display.Y >= di.rects.Display.Max.Y {
      // Done
      di.pt.Display.X -= 1
      di.pt.Display.Y -= 1
      return nil
    }
    di.pt.Display.X = di.rects.Display.Min.X
    di.pt.Work.X    = di.rects.Work.Min.X
    di.pt.Work.Y   -= di.dy

  } else {
    di.pt.Work.X  += di.dx
  }

  return &di.pt
}

// ------------------------------------------------------------------------------------------------------------

func MakeGrid(dleft, dtop, dright, dbottom int, left, top, right, bottom float64, strategy Strategy) *Both {

  fdleft, fdtop, fdright, fdbottom := float64(dleft), float64(dtop), float64(dright), float64(dbottom)

  origWidth     := right - left
  origHeight    := top - bottom

  unitsPerPixelX := (right - left) / (fdright - fdleft)
  unitsPerPixelY := (top - bottom) / (fdbottom - fdtop)

  fixX := func(unitsPerPixel float64) {
    unitsPerPixelX = unitsPerPixel
    newWidth := unitsPerPixel * (fdright - fdleft)
    diff := newWidth - origWidth
    half := diff / 2.0
    left, right = left-half, right+half
  }

  fixY := func(unitsPerPixel float64) {
    unitsPerPixelY = unitsPerPixel
    newHeight := unitsPerPixel * (fdbottom - fdtop)
    diff := newHeight - origHeight
    half := diff / 2.0
    top, bottom = top+half, bottom-half
  }

  if strategy == SeeAll {
    if unitsPerPixelX < unitsPerPixelY {
      fixX(unitsPerPixelY)
    } else {
      fixY(unitsPerPixelX)
    }
  } else if strategy == ZoomedInCropped {
    if unitsPerPixelX < unitsPerPixelY {
      fixY(unitsPerPixelX)
    } else {
      fixX(unitsPerPixelY)
    }
  }

  mainRect := BothRect{
    Display: Rect(dleft, dtop, dright, dbottom),
    Work:    WkRect(left, top, right, bottom),
  }
  return &Both{Main: mainRect}
}


func NormalizeDisplayRect(left, top, right, bottom int) (int,int,int,int)  {
  left, right = inOrderInt(left, right)
  top, bottom = inOrderInt(top, bottom)

  return left, top, right, bottom
}

func NormalizeWorkRect(left, top, right, bottom float64) (float64,float64,float64,float64)  {
  left, right = inOrder(left, right)
  top, bottom = inReverseOrder(top, bottom)

  return left, top, right, bottom
}

func inOrderInt(a,b int) (int,int) {
  if b < a {
    return b,a
  }
  return a,b
}

func inReverseOrderInt(a,b int) (int,int) {
  if b < a {
    return a,b
  }
  return b,a
}

func inOrder(a,b float64) (float64,float64) {
  if b < a {
    return b,a
  }
  return a,b
}

func inReverseOrder(a,b float64) (float64,float64) {
  if b < a {
    return a,b
  }
  return b,a
}

