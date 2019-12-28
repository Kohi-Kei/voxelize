package model

import "math"

//X is a X coordinate
type X float64

//Y is a Y coordinate
type Y float64

//Z is a Z coordinate
type Z float64

//Point is
type Point struct {
	X
	Y
	Z
	RGBA
}

func (p *Point) translate(x X, y Y, z Z) {
	p.X += x
	p.Y += y
	p.Z += z
}

func (p *Point) quantized(quantizeSize float64) *Point {
	qx := math.Trunc(float64(p.X) / quantizeSize)
	qy := math.Trunc(float64(p.Y) / quantizeSize)
	qz := math.Trunc(float64(p.Z) / quantizeSize)
	return &Point{X: X(qx), Y: Y(qy), Z: Z(qz)}
}

func (p *Point) scale(size float64) {
	p.X = p.X * X(size)
	p.Y = p.Y * Y(size)
	p.Z = p.Z * Z(size)

}

//Points is
type Points struct {
	points []*Point
}

//Add is
func (points *Points) Add(p *Point) {
	points.points = append(points.points, p)
}
func (points *Points) translate(x X, y Y, z Z) {
	for _, point := range points.points {
		point.translate(x, y, z)
	}
}

//NewPoint is
func NewPoint(x X, y Y, z Z) *Point {
	return &Point{X: X(x), Y: Y(y), Z: Z(z), RGBA: RGBA{}}
}

func (points Points) getPoints() []*Point {
	return points.points
}

func (points *Points) scale(size float64) {
	for _, p := range points.getPoints() {
		p.scale(size)
	}
}
