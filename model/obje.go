package model

import (
	"fmt"
	"log"
)

//Obje is
type Obje struct {
	Points
}

//MaxScaleSize is
const MaxScaleSize float64 = 1000.0

func (o *Obje) translate(x X, y Y, z Z) {
	for _, point := range o.getPoints() {
		point.translate(x, y, z)
	}
}

//ScalePoints is
func (o *Obje) ScalePoints(scaleSize float64) error {
	if scaleSize > MaxScaleSize {
		log.Fatal("scale size is too large: ", scaleSize)
	}
	o.scale(scaleSize)
	return nil
}

//CalcSize is
func (o *Obje) CalcSize() (x X, y Y, z Z) {
	minx, maxx := o.getMinMaxX()
	miny, maxy := o.getMinMaxY()
	minz, maxz := o.getMinMaxZ()
	//TODO: delete this print
	fmt.Println("x: ", maxx, minx)
	fmt.Println("y: ", maxy, miny)
	fmt.Println("z: ", maxz, minz)
	return X(maxx - minx), Y(maxy - miny), Z(maxz - minz)
}

func (o *Obje) calcMinCoordiane() (minx, miny, minz float64) {
	minx, _ = o.getMinMaxX()
	miny, _ = o.getMinMaxY()
	minz, _ = o.getMinMaxZ()
	return
}

func (o *Obje) getMinMaxX() (minx float64, maxx float64) {
	minx, maxx = 10e5, -10e5
	for _, point := range o.getPoints() {
		if float64(point.X) < minx {
			minx = float64(point.X)
		}

		if float64(point.X) > maxx {
			maxx = float64(point.X)
		}
	}
	return
}

func (o *Obje) getMinMaxY() (miny float64, maxy float64) {
	miny, maxy = 10e5, -10e5
	for _, point := range o.getPoints() {
		if float64(point.Y) < miny {
			miny = float64(point.Y)
		}

		if float64(point.Y) > maxy {
			maxy = float64(point.Y)
		}
	}
	return
}
func (o *Obje) getMinMaxZ() (minz float64, maxz float64) {
	minz, maxz = 10e5, -10e5
	for _, point := range o.getPoints() {
		if float64(point.Z) < minz {
			minz = float64(point.Z)
		}

		if float64(point.Z) > maxz {
			maxz = float64(point.Z)
		}
	}
	return
}
