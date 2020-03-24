package obj_test

import (
	"testing"

	"github.com/Kohi-Kei/voxelize/model"

	"github.com/Kohi-Kei/voxelize/exporter/obj"
)

// type Voxel struct {
// 	X ix `json:"x"int`
// 	Y iy `json:"y"int`
// 	Z iz `json:"z"int`
// 	RGBA
// }

// //Voxels is
// type Voxels []*Voxel

func TestMain(t *testing.T) {

	// voxels := model.Voxels{&model.Voxel{X: 1, Y: 1, Z: 1}}
	voxels := model.Voxels{
		&model.Voxel{X: 0, Y: 1, Z: 1, RGBA: model.RGBA{R: 255, G: 0, B: 0, A: 255}},
		&model.Voxel{X: 1, Y: 0, Z: 1, RGBA: model.RGBA{R: 255, G: 0, B: 0, A: 255}},
		&model.Voxel{X: 1, Y: 1, Z: 1, RGBA: model.RGBA{R: 255, G: 0, B: 0, A: 255}},
	}

	objExporter := obj.NewExporter(&voxels)
	voxels.Save(objExporter)
}
