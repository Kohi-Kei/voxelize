package main

import (
	"github.com/Kohi-Kei/voxelize/exporter/json"
	"github.com/Kohi-Kei/voxelize/exporter/obj"
	"github.com/Kohi-Kei/voxelize/loader"
	"github.com/Kohi-Kei/voxelize/model"
)

const voxelSize model.VoxelSize = model.VoxelSize(0.1)
const minPointsNum model.MinPointsNum = model.MinPointsNum(1)

func main() {

	fp := loader.OpenFilePointer()

	objeLoader := loader.New(loader.OBJ, fp)
	obje := objeLoader.Execute()

	voxelization := model.NewVoxelization(&obje, voxelSize)
	voxels := voxelization.CreateVoxels(minPointsNum)

	jsonExporter := json.NewJSONExporter(&voxels)
	voxels.Save(jsonExporter)

	objExporter := obj.NewExporter(&voxels)
	voxels.Save(objExporter)
}
