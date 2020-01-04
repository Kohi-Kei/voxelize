package obj

import (
	"os"

	"github.com/Kohi-Kei/voxelize/model"
)

//Exporter exports a .obj file
type Exporter struct {
	*lines
}

//NewNonColoreExporter returns a new Exporter
func NewNonColoreExporter(voxels *model.Voxels) *Exporter {
	lines := newLines()
	meshObje := model.NewMeshObje(voxels)
	lines.addGroup()
	lines.addVertices(meshObje)
	lines.addNoColoredFaces(meshObje)
	//now dont add normal vector
	return &Exporter{
		lines: lines,
	}
}

//Export write .obj files form voxel data
func (exporter *Exporter) Export() {
	//TODO: 外部ファイルに出力ファイルを書く
	fp, _ := os.Create("./asset/voxels.obj")
	defer fp.Close()
	exporter.write(fp)
}
