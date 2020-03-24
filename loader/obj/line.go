package obj

import (
	"strconv"
	"strings"

	"github.com/Kohi-Kei/voxelize/model"
	"github.com/Kohi-Kei/voxelize/model/obj"
)

type line struct {
	fields []string
}

func newLine(text string) *line {
	fields := strings.Fields(text)
	return &line{fields: fields}
}

func (line *line) fieldNum() int {
	return len(line.fields)
}

func (line *line) getField(i int) string {
	return line.fields[i]
}

func (line *line) geometryType() obj.GeometoryType {
	return obj.GeometoryType(line.getField(0))
}

func (line *line) toPoint() *model.Point {
	x, y, z := toCoordinates(line.getField(1), line.getField(2), line.getField(3))
	return model.NewPoint(x, y, z, model.RGBA{})
}

func toCoordinates(array ...string) (model.X, model.Y, model.Z) {
	x, errx := strconv.ParseFloat(array[0], 64)
	y, erry := strconv.ParseFloat(array[1], 64)
	z, errz := strconv.ParseFloat(array[2], 64)
	if errx != nil {
		panic(errx)
	}
	if erry != nil {
		panic(erry)
	}
	if errz != nil {
		panic(errz)
	}
	return model.X(x), model.Y(y), model.Z(z)
}
