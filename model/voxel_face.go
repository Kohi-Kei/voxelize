package model

//TriangleNum is number of voxel edges
const TriangleNum int = 12

type voxelID int

//Face is face of voxel
type Face struct {
	Points  [3]*pointID
	VoxelID voxelID
}

//Faces is
type Faces []*Face

func (faces *Faces) merge(id voxelID, childs *Faces) {
	for triIndex, child := range *childs {
		(*faces)[TriangleNum*int(id)+triIndex] = child
	}
}
