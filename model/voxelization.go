package model

//MinPointsNum is
type MinPointsNum uint32

//Voxelization is
type Voxelization struct {
	lattice *lattice
}

// func (voxelization *Voxelization) execute() {
// 	lattice := voxelization.lattice
// 	lattice.pointsToCells()
// }

//NewVoxelization is
func NewVoxelization(obje *Obje, voxelSize VoxelSize) Voxelization {
	lattice := newLattice(obje, voxelSize)
	return Voxelization{lattice}
}

//CreateVoxels is
func (voxelization *Voxelization) CreateVoxels(minPointsNum MinPointsNum) Voxels {
	lattice := voxelization.lattice
	lattice.pointsToCells()
	return lattice.toVoxels(int(minPointsNum))
}
