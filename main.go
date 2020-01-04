package main

import (
	"io/ioutil"
	"os"

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

	ioutil.WriteFile("./asset/voxels.json", voxels.ToJSONBytes(), os.ModePerm)
	objExporter := obj.NewNonColoreExporter(&voxels)
	voxels.Save(objExporter)
}
