package json

import (
	"io/ioutil"
	"os"

	"github.com/Kohi-Kei/voxelize/model"
)

type bytes []byte

//Exporter exports a json file
type Exporter struct {
	bytes
}

//NewJSONExporter returns a new Exporter
func NewJSONExporter(voxels *model.Voxels) *Exporter {
	return &Exporter{
		bytes: bytes(voxels.ToJSONBytes()),
	}
}

//Export write .obj files form voxel data
func (exporter *Exporter) Export() {
	//TODO: 出力ファイルの外出し
	ioutil.WriteFile("./asset/voxels.json", exporter.bytes, os.ModePerm)
}
