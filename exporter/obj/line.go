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

//TODO: //0のところに法線ベクトルを入れなきゃだめ
func newNonColoredFace(face *model.Face) *line {
	return &line{
		fields: fields{
			fmt.Sprintf("%d//0", int(*face.Points[0])),
			fmt.Sprintf("%d//0", int(*face.Points[1])),
			fmt.Sprintf("%d//0", int(*face.Points[2]))},
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

func (lines *lines) addNoColoredFaces(meshObje *model.MeshObje) {
	for _, face := range meshObje.Faces {
		faceLine := newNonColoredFace(face)
		*lines = append(*lines, faceLine)
	}
}

func (lines *lines) write(fp *os.File) {
	for _, line := range *lines {
		fmt.Fprintln(fp, line.connectFields())
	}
}
