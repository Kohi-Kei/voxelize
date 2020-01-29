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
	var originFaces model.OriginFaces
	var originVertexTexs model.OriginVertexTexs

	// definite 'scanner' for the file
	scanner := bufio.NewScanner(loader.fp)
	// scan all lines of the file
	for scanner.Scan() {
		// get a line from the file
		line := newLine(scanner.Text())
		// Insert some info. (vertex,face,texture)
		if line.fieldNum() != 0 {

			switch line.geometryType() {
			// insert Vertex info.
			case obj.Vertex:
				points.Add(line.toPoint())

			// insert Face info.
			case obj.Face:
				originFaces.Add(line.toFace())

			// insert Texture Info
			case obj.TextureVector:
				originVertexTexs.Add(line.toVertexTex())

			default:
			}
		}
	}
	return model.Obje{Points: points}
}
