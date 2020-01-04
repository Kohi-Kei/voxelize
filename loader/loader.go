package loader

import (
	"os"

	"github.com/Kohi-Kei/voxelize/adapter"
	"github.com/Kohi-Kei/voxelize/loader/obj"
)

// FormatType is the file type you'd like to laod
type FormatType string

// OBJ is .obj formattt file
const OBJ FormatType = "obj"

// New is
func New(format FormatType, fp *os.File) adapter.Loader {
	switch format {
	case OBJ:
		return obj.NewLoader(fp)
	default:
	}
	panic("Unsupported format type")
}

//OpenFilePointer make a file pointer to a 3d file
func OpenFilePointer() *os.File {
	var fp *os.File
	var err error

	if len(os.Args) < 2 {
		panic("please type input file.")
	} else {
		fp, err = os.Open(os.Args[1])
		if err != nil {
			panic(err)
		}
	}
	return fp
}
