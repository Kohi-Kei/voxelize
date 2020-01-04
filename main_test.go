package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"testing"

	"github.com/Kohi-Kei/voxelize/loader"
	"github.com/Kohi-Kei/voxelize/model"
	"github.com/Kohi-Kei/voxelize/writer"
)

func TestMain(t *testing.T) {

	fp, _ := os.Open("./asset/texturedMesh.obj")

	objeLoader := loader.New(loader.OBJ, fp)
	obje := objeLoader.Execute()

	voxelization := model.NewVoxelization(&obje, voxelSize)
	voxels := voxelization.CreateVoxels(minPointsNum)

	// ioutil.WriteFile("./asset/voxels.json", voxels.ToJSONBytes(), os.ModePerm)
	writer.Writer(&voxels)
}
func TestObjWriter(t *testing.T) {
	voxels := model.Voxels{&model.Voxel{X: 1, Y: 1, Z: 1}}
	writer.Writer(&voxels)
}

func TestMtlFile(t *testing.T) {
	// 1×1 ピクセル
	img := image.NewNRGBA(image.Rect(0, 0, 2, 1))
	img.Set(0, 0, color.RGBA{255, 0, 0, 255})
	img.Set(1, 0, color.RGBA{0, 255, 0, 255})
	file, _ := os.Create("./asset/blue_texture.png")
	defer file.Close()

	png.Encode(file, img)
}
