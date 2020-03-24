package model

//MeshObje has 8 points and 12 triangle faces
type MeshObje struct {
	Points
	Faces
}

//NewMeshObje convert objData from voxels
//A voxel point turns into 8 corner points
func NewMeshObje(voxels *Voxels) *MeshObje {
	corners := make([]*Point, voxels.GetCornerPointNum())
	faces := make(Faces, voxels.GetTriangleFaceNum())
	triangulateCorners(voxels, corners, faces)
	return &MeshObje{
		Points: Points(corners),
		Faces:  faces}
}

func triangulateCorners(voxels *Voxels, objPoints []*Point, objFaces Faces) {
	for voxelIndex, voxel := range *voxels {
		corners := calcCorners(newCenterPoint(voxel))

		setCorners(objPoints, voxelIndex, corners)
		triangulate(objFaces, voxelID(voxelIndex), corners)
	}
}

func setCorners(objPoints []*Point, voxelIndex int, corners Corners) {
	for index, corner := range corners {
		objPoints[CornerPointNum*voxelIndex+index] = corner.point
	}
}

func triangulate(objFaces Faces, voxelID voxelID, corners Corners) {
	tries := corners.triangulate(voxelID)
	objFaces.merge(voxelID, &tries)
}
