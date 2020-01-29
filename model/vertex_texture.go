package model

// TexX is x-coordinate of vertex on a texrure image
type TexX float64

// TexY is x-coordinate of vertex on a texrure image
type TexY float64

//OriginVertexTex is
type OriginVertexTex struct {
	TexX
	TexY
}

//OriginVertexTexs is
type OriginVertexTexs struct {
	originVertexTexs []*OriginVertexTex
}

//Add is
func (originVertexTexs *OriginVertexTexs) Add(vt *OriginVertexTex) {
	originVertexTexs.originVertexTexs = append(originVertexTexs.originVertexTexs, vt)
}

/*
func (points *Points) translate(x X, y Y, z Z) {
	for _, point := range points.points {
		point.translate(x, y, z)
	}
}*/

//NewVertexTex is
func NewVertexTex(x TexX, y TexY) *OriginVertexTex {
	return &OriginVertexTex{TexX: TexX(x), TexY: TexY(y)}
}

func (originVertexTexs OriginVertexTexs) getOriginVertexTexs() []*OriginVertexTex {
	return originVertexTexs.originVertexTexs
}

/*
func (VertexTexs *VertexTexs) scale(size float64) {
	for _, f := range VertexTexs.getVertexTexs() {
		f.scale(size)
	}
}
*/
