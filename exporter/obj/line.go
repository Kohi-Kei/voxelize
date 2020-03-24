package obj

import (
	"bytes"
	"fmt"
	"os"

	"github.com/Kohi-Kei/voxelize/model"
	"github.com/Kohi-Kei/voxelize/model/obj"
)

type fields []string

type line struct {
	fields
	geometory obj.GeometoryType
}

type lines []*line

func newLines() *lines {
	return &lines{}
}

func newGroup() *line {
	return &line{
		fields:    fields{"voxelize"},
		geometory: obj.Group,
	}
}

func newVertex(point *model.Point) *line {
	return &line{
		fields: fields{
			fmt.Sprintf("%f", float64(point.X)),
			fmt.Sprintf("%f", float64(point.Y)),
			fmt.Sprintf("%f", float64(point.Z))},
		geometory: obj.Vertex}
}

func newTextureVector(id int, pointNum int) *line {
	return &line{
		fields: fields{
			fmt.Sprintf("%.8f", float64(0.0)),
			fmt.Sprintf("%.8f", float64(id)/float64(pointNum))},
		geometory: obj.TextureVector}
}

//TODO: //0のところに法線ベクトルを入れなきゃだめ
func newNonColoredFace(face *model.Face) *line {
	return &line{
		fields: fields{
			fmt.Sprintf("%d//0", int(*face.Points[0])),
			fmt.Sprintf("%d//0", int(*face.Points[1])),
			fmt.Sprintf("%d//0", int(*face.Points[2]))},
		geometory: obj.Face}
}

//TODO: //0のところに法線ベクトルを入れなきゃだめ
func newColoredFace(face *model.Face, faceIndex int) *line {
	return &line{
		fields: fields{
			fmt.Sprintf("%d/%d", int(*face.Points[0]), faceIndex),
			fmt.Sprintf("%d/%d", int(*face.Points[1]), faceIndex),
			fmt.Sprintf("%d/%d", int(*face.Points[2]), faceIndex)},
		geometory: obj.Face}
}

func (line *line) connectFields() string {
	var buffer bytes.Buffer
	buffer.WriteString(string(line.geometory))
	for _, field := range line.fields {
		buffer.WriteString(" ")
		buffer.WriteString(field)
	}
	return buffer.String()
}

func (lines *lines) addGroup() {
	*lines = append(*lines, newGroup())
}

func (lines *lines) addVertices(meshObje *model.MeshObje) {
	for _, point := range meshObje.Points {
		vertexLine := newVertex(point)
		*lines = append(*lines, vertexLine)
	}
}

func (lines *lines) addTextureVector(meshObje *model.MeshObje) {
	pointNum := len(meshObje.Points)
	// 8 points has the same color,  so skip index
	for i := 0; i < len(meshObje.Points); i = i + model.CornerPointNum {
		vtLine := newTextureVector(i, pointNum)
		*lines = append(*lines, vtLine)
	}
}

func (lines *lines) addNoColoredFaces(meshObje *model.MeshObje) {
	for _, face := range meshObje.Faces {
		//newNonColoredFaceを治すこと//
		faceLine := newNonColoredFace(face)
		*lines = append(*lines, faceLine)
	}
}

func (lines *lines) addColoredFaces(meshObje *model.MeshObje) {
	for i := 0; i < len(meshObje.Faces); i++ {
		faceLine := newColoredFace(meshObje.Faces[i], i/model.TriangleNum)
		*lines = append(*lines, faceLine)
	}
}

func (lines *lines) write(fp *os.File) {
	for _, line := range *lines {
		fmt.Fprintln(fp, line.connectFields())
	}
}
