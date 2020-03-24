package model

import (
	"encoding/json"
	"log"

	"github.com/Kohi-Kei/voxelize/exporter"
)

//i means integer
type ix int
type iy int
type iz int

//Voxel is voxel
type Voxel struct {
	X ix `json:"x"int`
	Y iy `json:"y"int`
	Z iz `json:"z"int`
	RGBA
}

//Voxels is
type Voxels []*Voxel

func newVoxel(ix ix, iy iy, iz iz, rgba RGBA) *Voxel {
	return &Voxel{ix, iy, iz, RGBA(rgba)}
}

// ToJSONBytes is
func (voxels *Voxels) ToJSONBytes() []byte {
	bytes, ok := json.Marshal(voxels)
	if ok != nil {
		log.Fatal(ok)
	}
	return bytes
}

// GetColors returns voxel color
func (voxels *Voxels) GetColors() Colors {
	var colors []RGBA
	for _, voxel := range *voxels {
		colors = append(colors, voxel.RGBA)
	}
	return Colors(colors)
}

// GetVoxelNum is the number of voxels
func (voxels *Voxels) GetVoxelNum() int {
	return len(*voxels)
}

// GetCornerPointNum is
func (voxels *Voxels) GetCornerPointNum() int {
	return voxels.GetVoxelNum() * CornerPointNum
}

// GetTriangleFaceNum is
func (voxels *Voxels) GetTriangleFaceNum() int {
	return voxels.GetVoxelNum() * TriangleNum
}

//Save .obj file
func (voxels *Voxels) Save(exporter exporter.Exporter) {
	exporter.Export()
}
