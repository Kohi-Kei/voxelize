package obj

import "github.com/Kohi-Kei/voxelize/model/obj"

type fields []string

type line struct {
	fields
	obj.GeometoryType
}

type lines []*line
