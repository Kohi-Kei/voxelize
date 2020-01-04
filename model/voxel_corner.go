package model

type pointID int

func newPointID(id int, voxelID voxelID) *pointID {
	pointID := pointID(id + int(voxelID)*CornerPointNum)
	return &pointID
}

//CornerPointNum is number of voxel points
const CornerPointNum int = 8

//VoxelHalfLength is lenght of voxel edge
const VoxelHalfLength float32 = 0.5

//Corner is a voxel corner point
type Corner struct {
	point *Point
}

//Corners are voxel corner points
type Corners [CornerPointNum]*Corner

func calcCorners(point *centerPoint) Corners {

	c1 := &Corner{
		point: point.TranslatedNewPoint(-X(VoxelHalfLength), +Y(VoxelHalfLength), Z(VoxelHalfLength))}
	c2 := &Corner{
		point: point.TranslatedNewPoint(-X(VoxelHalfLength), +Y(VoxelHalfLength), -Z(VoxelHalfLength))}
	c3 := &Corner{
		point: point.TranslatedNewPoint(+X(VoxelHalfLength), +Y(VoxelHalfLength), +Z(VoxelHalfLength))}
	c4 := &Corner{
		point: point.TranslatedNewPoint(+X(VoxelHalfLength), +Y(VoxelHalfLength), -Z(VoxelHalfLength))}

	c5 := &Corner{
		point: point.TranslatedNewPoint(-X(VoxelHalfLength), -Y(VoxelHalfLength), +Z(VoxelHalfLength))}
	c6 := &Corner{
		point: point.TranslatedNewPoint(-X(VoxelHalfLength), -Y(VoxelHalfLength), -Z(VoxelHalfLength))}
	c7 := &Corner{
		point: point.TranslatedNewPoint(+X(VoxelHalfLength), -Y(VoxelHalfLength), +Z(VoxelHalfLength))}
	c8 := &Corner{
		point: point.TranslatedNewPoint(+X(VoxelHalfLength), -Y(VoxelHalfLength), -Z(VoxelHalfLength))}

	return Corners{c1, c2, c3, c4, c5, c6, c7, c8}
}

func (Corners *Corners) selectThreePoints(v1, v2, v3 int, voxelID voxelID) *Face {
	return &Face{
		Points: [3]*pointID{
			newPointID(v1, voxelID),
			newPointID(v2, voxelID),
			newPointID(v3, voxelID)},
		VoxelID: voxelID}
}

func (Corners *Corners) triangulate(id voxelID) Faces {

	//3つの数字にそれぞれ位置を足す?
	tri0 := Corners.selectThreePoints(1, 2, 3, id)
	tri1 := Corners.selectThreePoints(3, 2, 4, id)

	tri2 := Corners.selectThreePoints(2, 6, 4, id)
	tri3 := Corners.selectThreePoints(4, 6, 8, id)

	tri4 := Corners.selectThreePoints(1, 5, 2, id)
	tri5 := Corners.selectThreePoints(2, 5, 6, id)

	tri6 := Corners.selectThreePoints(4, 8, 3, id)
	tri7 := Corners.selectThreePoints(3, 8, 7, id)

	tri8 := Corners.selectThreePoints(3, 7, 1, id)
	tri9 := Corners.selectThreePoints(1, 7, 5, id)

	tri10 := Corners.selectThreePoints(6, 5, 8, id)
	tri11 := Corners.selectThreePoints(8, 5, 7, id)

	return Faces([]*Face{tri0, tri1, tri2, tri3, tri4,
		tri5, tri6, tri7, tri8, tri9, tri10, tri11})
}
