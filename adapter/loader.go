package adapter

import (
	"github.com/Kohi-Kei/voxelize/model"
)
//Loader is a interface of  3D model file loaders
type Loader interface{
	Execute() model.Obje
}
