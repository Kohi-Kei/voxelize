package obj

import (
	"bufio"
	"os"

	"github.com/Kohi-Kei/voxelize/model"
	"github.com/Kohi-Kei/voxelize/model/obj"
)

//Loader is "obj format" loader
type Loader struct {
	fp *os.File
}

//NewLoader return new Obj Loader Object
func NewLoader(fp *os.File) *Loader {
	return &Loader{fp: fp}
}

//Execute is
func (loader *Loader) Execute() model.Obje {
	var points model.Points
	scanner := bufio.NewScanner(loader.fp)
	for scanner.Scan() {
		line := newLine(scanner.Text())
		if line.fieldNum() != 0 {
			switch line.geometryType() {
			case obj.Vertex:
				points.Add(line.toPoint())
			case obj.Face:
			default:
			}
		}
	}
	return model.Obje{Points: points}
}
