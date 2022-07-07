package corkscrew

import (
  "image"
)

/* Copyright © 2022 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

// -------------------------------------------------------------------------------------------------------------------

func NormalizeRealm(realmRect Rec2, rec image.Rectangle) Rec2 {
  //realmRect: Rec2{Min: Vec2{-2.1, 1.2}, Max: Vec2{1.2, -1.2}},
  //rec: image.Rect(0, 0, 1200, 800),
  //return Rec2{Vec2{-2.25, 1.2}, Vec2{1.35, -1.2}}

  x0,y0,x1,y1 := realmRect.Parts()

  xUnitsPerPixel := /* 3.3 / 1200  -->  0.00275 */ (realmRect.Max.X - realmRect.Min.X) / float32(rec.Dx())
  yUnitsPerPixel := /* 2.4 / 800   -->  0.003 */   (realmRect.Min.Y - realmRect.Max.Y) / float32(rec.Dy())

  if yUnitsPerPixel > xUnitsPerPixel {
    // Change X
    xUnitsPerPixel = yUnitsPerPixel

    newSpan := /* 0.003 * 1200 */   xUnitsPerPixel * float32(rec.Dx())    /* 3.6 */
    diff    := /* 3.6 - 3.3 */      newSpan - realmRect.Dx()              /* 0.3 */
    half    :=                      diff / 2.0                            /* 0.15 */

    return R2(x0 - half, y0, x1 + half, y1)
  } else {
    // Change Y
    yUnitsPerPixel = xUnitsPerPixel

    newSpan := /* 0.003 * 1200 */   yUnitsPerPixel * float32(rec.Dy())
    diff    := /* 3.6 - 3.3 */      newSpan - realmRect.Dy()
    half    :=                      diff / 2.0

    return R2(x0, y0 + half, x1, y1 - half)
  }
}

// -------------------------------------------------------------------------------------------------------------------

//RealmToScreen will return the Point the represents the input Vec2 point
func RealmToScreen(pt Vec2, realmRect Rec2, rec image.Rectangle) image.Point {
  //pt: Vec2{-1.0,1.0},
  //realmRect: Rec2{Min: Vec2{-2.25,1.2}, Max: Vec2{1.35,-1.2}},
  //rec: image.Rect(0, 0, 1200, 800),
  //return image.Point{X: 416 Y: 66}

  //      pt.X       pXdist     percentX       x
  // x:   -1.0  -->   1.25  -->  0.3472  -->  416

  ptXdist   := pt.X - realmRect.Min.X
  percentX  := ptXdist / realmRect.Dx()
  x         := int(percentX * float32(rec.Dx()))

  ptYdist   := pt.Y - realmRect.Max.Y
  percentY  := 1.0 - (ptYdist / realmRect.Dy())
  y         := int(percentY * float32(rec.Dy()))

  return image.Point{X: x, Y: y}
}

// -------------------------------------------------------------------------------------------------------------------

//RealmToScreenVec2 will return the Point the represents the input Vec2 point
func RealmToScreenVec2(pt Vec2, realmRectMin, realmRectMax Vec2, rec image.Rectangle) image.Point {
  realmRect := Rec2{realmRectMin, realmRectMax}
  return RealmToScreen(pt, realmRect, rec)
}

// -------------------------------------------------------------------------------------------------------------------

func ScreenToRealm(pt image.Point, realmRect Rec2, rec image.Rectangle) Vec2 {
  //pt: image.Point{X: 416 Y: 66},
  //realmRect: Rec2{Min: Vec2{-2.25,1.2}, Max: Vec2{1.35,-1.2}},
  //rec: image.Rect(0, 0, 1200, 800),
  //return Vec2{-1.0,1.0}

  //       x      percentX     pXdist        x
  // x:  416  -->  0.3472  -->  1.25  -->  -1.o

  percentX := float32(pt.X) / float32(rec.Dx())
  pXdist   := percentX * realmRect.Dx()
  x        := pXdist + realmRect.Min.X

  percentY := 1.0 - (float32(pt.Y) / float32(rec.Dy()))
  pYdist   := percentY * realmRect.Dy()
  y        := pYdist + realmRect.Max.Y

  return Vec2{x, y}
}

// -------------------------------------------------------------------------------------------------------------------

func ScreenToRealmVec2(pt image.Point, realmRectMin, realmRectMax Vec2, rec image.Rectangle) Vec2 {
  realmRect := Rec2{realmRectMin, realmRectMax}
  return ScreenToRealm(pt, realmRect, rec)
}
