package model

// vertexNum is a number of vertex
type vertexNum int32
type vertexNumArray []*vertexNum

// texNum is a number of vertex
type texNum int32
type texNumArray []*texNum

//OriginFace is
type OriginFace struct {
	vertexNumArray
	texNumArray
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
type OriginFaces struct {
	originfaces []*OriginFace
}

//Add is
func (originfaces *OriginFaces) Add(f *OriginFace) {
	originfaces.originfaces = append(originfaces.originfaces, f)
}

/*
func (points *Points) translate(x X, y Y, z Z) {
	for _, point := range points.points {
		point.translate(x, y, z)
	}
}*/

//NewFace is
func NewFace(vertexArray vertexNumArray, textureArray texNumArray) *OriginFace {
	return &OriginFace{vertexNumArray: vertexNumArray(vertexArray), texNumArray: texNumArray(textureArray)}
}

func (originfaces OriginFaces) getOriginFaces() []*OriginFace {
	return originfaces.originfaces
}

/*
func (faces *Faces) scale(size float64) {
	for _, f := range faces.getFaces() {
		f.scale(size)
	}
}
*/
