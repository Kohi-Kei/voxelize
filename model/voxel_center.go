package model

//voxelCenterPoint is the center point of voxel
type centerPoint struct {
	Point
}

type cornerPoint struct {
	Point
}

func (centerPoint *centerPoint) NewCornerPoint(dx X, dy Y, dz Z) *cornerPoint {
	point := *NewPoint(centerPoint.X+dx, centerPoint.Y+dy, centerPoint.Z+dz)
	return &cornerPoint{point}
}

//NewVoxelCenterPoint return voxel center point
func newCenterPoint(voxel *Voxel) *centerPoint {
	return &centerPoint{
		Point{
			X: X(voxel.X),
			Y: Y(voxel.Y),
			Z: Z(voxel.Z),
		},
	}
}
