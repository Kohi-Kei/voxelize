package model

//VoxelSize is
type VoxelSize float64

type lattice struct {
	VoxelSize
	quantizedObje *Obje
	cells
	xsize int
	ysize int
	zsize int
}

func (lattice *lattice) cell(xi int, yi, int, zi int) *cell {
	return lattice.cells.at(X(xi), Y(yi), Z(zi))
}

func (lattice *lattice) pointsToCells() {
	cells := lattice.cells
	for _, point := range lattice.quantizedObje.getPoints() {
		cells.to(point)
	}
}

func (lattice *lattice) toVoxels(minPointsNum int) Voxels {
	var voxels []*Voxel
	for xi := 0; xi < lattice.xsize; xi++ {
		for yi := 0; yi < lattice.ysize; yi++ {
			for zi := 0; zi < lattice.zsize; zi++ {
				if lattice.cells.pointsNum(X(xi), Y(yi), Z(zi)) >= minPointsNum {
					voxels = append(voxels, newVoxel(ix(xi), iy(yi), iz(zi)))
				}
			}
		}
	}
	return voxels
}

func newLattice(obje *Obje, voxelSize VoxelSize) *lattice {
	obje.scale(1.0 / float64(voxelSize))
	quantizedObje := newQuantizedObje(obje, float64(voxelSize))
	xsize, ysize, zsize := quantizedObje.CalcSize()
	minx, miny, minz := quantizedObje.calcMinCoordiane()
	quantizedObje.translate(X(-minx), Y(-miny), Z(-minz))
	return &lattice{
		VoxelSize:     voxelSize,
		quantizedObje: quantizedObje,
		cells:         newCells(quantizedObje),
		xsize:         int(xsize) + 1,
		ysize:         int(ysize) + 1,
		zsize:         int(zsize) + 1,
	}
}

func newQuantizedObje(obje *Obje, voxelSize float64) *Obje {
	originalPoints := obje.getPoints()
	var quantizedPoints []*Point
	for _, originalPoint := range originalPoints {
		quantizedPoint := originalPoint.quantized(voxelSize)
		quantizedPoints = append(quantizedPoints, quantizedPoint)
	}
	return &Obje{
		Points: Points(quantizedPoints)}
}
