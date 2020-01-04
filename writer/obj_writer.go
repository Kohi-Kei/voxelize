package writer

import (
	"fmt"
	"os"

	"github.com/Kohi-Kei/voxelize/model"
)

// type writer struct {
// 	voxels *model.Voxels
// 	fp     *os.File
// }

//Writer is
func Writer(voxels *model.Voxels) {
	objData := model.NewMeshObje(voxels)
	fp, _ := os.Create("./asset/voxels.obj")
	defer fp.Close()

	fmt.Fprintf(fp, "g voxelize\n")

	points := objData.Points
	for _, point := range points {
		fmt.Fprintf(fp, "v %g %g %g\n", point.X, point.Y, point.Z)
	}
	fmt.Fprint(fp, "vn -1 0 0\n")
	fmt.Fprint(fp, "vn 1 0 0\n")
	fmt.Fprint(fp, "vn 0 -1 0\n")
	fmt.Fprint(fp, "vn 0 1 1\n")
	fmt.Fprint(fp, "vn 0 0 -1\n")
	fmt.Fprint(fp, "vn 0 0 1\n")
	for _, triangle := range objData.Faces {
		fmt.Fprintf(fp, "f %d//%d %d//%d %d//%d\n", *triangle.Points[0], triangle.VoxelID, *triangle.Points[1], triangle.VoxelID, *triangle.Points[2], triangle.VoxelID)
	}
}
