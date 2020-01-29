package model

// vertexNum is a number of vertex
type VertexNum int32
type VertexNumArray []*VertexNum

// texNum is a number of vertex
type TexNum int32
type TexNumArray []*TexNum

//OriginFace is
type OriginFace struct {
	VertexNumArray
	TexNumArray
}

/*
func (p *Point) translate(x X, y Y, z Z) {
	p.X += x
	p.Y += y

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
func NewFace(vertexArray []VertexNum, textureArray []TexNum) *OriginFace {
	var vertexArrayPointer VertexNumArray
	var texNumArrayPointer TexNumArray

	for index := range vertexArray {
		vertexArrayPointer = append(vertexArrayPointer, &vertexArray[index])
	}

	for index := range textureArray {
		texNumArrayPointer = append(texNumArrayPointer, &textureArray[index])
	}

	return &OriginFace{VertexNumArray: vertexArrayPointer, TexNumArray: texNumArrayPointer}
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
