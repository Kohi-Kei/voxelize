package model

import (
	"encoding/json"
	"log"
)

type ix int
type iy int
type iz int

type voxel struct {
	X ix `json:"x"int`
	Y iy `json:"y"int`
	Z iz `json:"z"int`
	RGBA
}

//Voxels is
type Voxels []*voxel

//TODO
func newVoxel(ix ix, iy iy, iz iz) *voxel {
	return &voxel{ix, iy, iz, RGBA{}}
}

// ToJSONBytes is
func (voxels *Voxels) ToJSONBytes() []byte {
	bytes, ok := json.Marshal(voxels)
	if ok != nil {
		log.Fatal(ok)
	}
	return bytes
}
