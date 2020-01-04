package model

//D is a A coordinate
type D float64

//E is a Y coordinate
type E float64

//C is a Z coordinate
type C float64

//Face is
type Face struct {
	D
	E
	C
	RGBA
}

/*
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
*/
//Faces is
type Faces struct {
	faces []*Face
}

//Add is
func (faces *Faces) Add(f *Face) {
	faces.faces = append(faces.faces, f)
}

/*
func (points *Points) translate(x X, y Y, z Z) {
	for _, point := range points.points {
		point.translate(x, y, z)
	}
}*/

//NewFace is
func NewFace(x D, y E, z C) *Face {
	return &Face{D: D(x), E: E(y), C: C(z), RGBA: RGBA{}}
}

func (faces Faces) getFaces() []*Face {
	return faces.faces
}

/*
func (faces *Faces) scale(size float64) {
	for _, f := range faces.getFaces() {
		f.scale(size)
	}
}
*/
