package both

/* Copyright © 2022 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

type Strategy int
const (
  SeeAll  Strategy = iota
  ZoomedInCropped
  DoNotChange
)
func (s Strategy) String() string {
  return [...]string{"SeeAll", "ZoomedInCropped", "DoNotChange"}[s]
}

// -----------------------------------------------------------------------------------------------------------------

var (
  unique int = 0
)

// -----------------------------------------------------------------------------------------------------------------

type Both struct {
  Id      int
  Main    BothRect        /* The rect for the whole display */
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
  x, y := di.pt.Display.X, di.pt.Display.Y

  x += 1
  if x >= di.rects.Display.Max.X {
    y += 1
    if y >= di.rects.Display.Max.Y {
      // Done
      return nil
    }
    di.pt.Display.X = di.rects.Display.Min.X
    di.pt.Display.Y = y
    di.pt.Work.X    = di.rects.Work.Min.X
    di.pt.Work.Y   -= di.dy

  } else {
    di.pt.Display.X    = x
    di.pt.Work.X      += di.dx
  }

  return &di.pt
}

// ------------------------------------------------------------------------------------------------------------

func (b *Both) Split() (*Both, *Both) {
  myWorkRect := b.Main.Work
  if b.Main.Work.Dx() > b.Main.Work.Dy() {
    dleft, dright := b.Main.Display.SplitOnX()
    //dright.Min.X = b.Main.Display.MidpointX()
    //dleft.Max.X = dright.Min.X - 1

    left := WkRect(myWorkRect.Min.X, myWorkRect.Min.Y, b.WorkPtFor(dleft.Max).X, myWorkRect.Max.Y)
    right := WkRect(b.WorkPtFor(dright.Min).X, myWorkRect.Min.Y, myWorkRect.Max.X, myWorkRect.Max.Y)
    return MakeGrid(dleft, left, DoNotChange), MakeGrid(dright, right, DoNotChange)

  } else {
    dtop, dbottom := b.Main.Display.SplitOnY()
    //dbottom.Min.Y = b.Main.Display.MidpointY()
    //dtop.Max.Y = dbottom.Min.Y

    top := WkRect(myWorkRect.Min.X, myWorkRect.Min.Y, myWorkRect.Max.X, b.WorkPtFor(dtop.Max).Y)
    bottom := WkRect(myWorkRect.Min.X, b.WorkPtFor(dbottom.Min).Y, myWorkRect.Max.X, myWorkRect.Max.Y)
    return MakeGrid(dtop, top, DoNotChange), MakeGrid(dbottom, bottom, DoNotChange)
  }

}

// ------------------------------------------------------------------------------------------------------------

func MakeGrid(/*id int,*/ drect DisplayRect, rect WorkRect, strategy Strategy) *Both {
  return MakeGrid8(/*id,*/ drect.Min.X, drect.Min.Y, drect.Max.X, drect.Max.Y, rect.Min.X, rect.Min.Y, rect.Max.X, rect.Max.Y, strategy)
}

// ------------------------------------------------------------------------------------------------------------

func MakeGrid8(/*id int,*/ dleft, dtop, dright, dbottom int, left, top, right, bottom float64, strategy Strategy) *Both {

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

  unique += 1
  return &Both{Id: unique, Main: mainRect}
}

// -----------------------------------------------------------------------------------------------------------------

func NormalizeDisplayRect(left, top, right, bottom int) (int,int,int,int)  {
  left, right = inOrderInt(left, right)
  top, bottom = inOrderInt(top, bottom)

  return left, top, right, bottom
}

// -----------------------------------------------------------------------------------------------------------------

func NormalizeWorkRect(left, top, right, bottom float64) (float64,float64,float64,float64)  {
  left, right = inOrder(left, right)
  top, bottom = inReverseOrder(top, bottom)

  return left, top, right, bottom
}

// -----------------------------------------------------------------------------------------------------------------

func inOrderInt(a,b int) (int,int) {
  if b < a {
    return b,a
  }
  return a,b
}

// -----------------------------------------------------------------------------------------------------------------

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

// -----------------------------------------------------------------------------------------------------------------

func inReverseOrder(a,b float64) (float64,float64) {
  if b < a {
    return a,b
  }
  return b,a
}

