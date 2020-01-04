package obj

//GeometoryType is type of Geometory (ex. v:Vertex, f: Face)
type GeometoryType string

//Vertex is
const Vertex GeometoryType = GeometoryType("v")

//Face is
const Face GeometoryType = GeometoryType("f")

//NornalVector is
const NornalVector GeometoryType = GeometoryType("vn")

//Group is
const Group GeometoryType = GeometoryType("g")

//TextureVector represents coordinates of texture images
const TextureVector GeometoryType = GeometoryType("vt")
