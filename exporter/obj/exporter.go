package obj

import (
	"image"
	"image/color"
	"image/png"
	"os"

	"github.com/Kohi-Kei/voxelize/model"
)

//Exporter exports a .obj file
type Exporter struct {
	*lines
	model.Colors
}

//NewNonColorExporter returns a new Exporter
func NewNonColorExporter(voxels *model.Voxels) *Exporter {
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

//NewExporter returns a initialized Exporter
func NewExporter(voxels *model.Voxels) *Exporter {
	lines := newLines()
	meshObje := model.NewMeshObje(voxels)
	lines.addGroup()
	lines.addVertices(meshObje)
	lines.addTextureVector(meshObje)
	lines.addColoredFaces(meshObje)
	//now dont add normal vector
	return &Exporter{
		lines:  lines,
		Colors: voxels.GetColors(),
	}
}

//Export write .obj files form voxel data
func (exporter *Exporter) Export() {
	exporter.exportTexturePixels()
	//TODO: 外部ファイルに出力ファイルを書く
	fp, _ := os.Create("./voxels.obj")
	defer fp.Close()
	exporter.write(fp)
}

func (exporter *Exporter) exportTexturePixels() {
	//Rect(width, height)
	img := image.NewNRGBA(image.Rect(0, 0, len(exporter.Colors), 1))
	for i, rgba := range exporter.Colors {
		img.Set(i, 0, color.RGBA{uint8(rgba.R), uint8(rgba.G), uint8(rgba.B), uint8(rgba.A)})
	}
	file, _ := os.Create("./texture.png")
	defer file.Close()
	png.Encode(file, img)
}
