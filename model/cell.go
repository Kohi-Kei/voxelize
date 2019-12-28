package model


type cell struct {
	points Points
}

func (c *cell) getPointNum() int {
	return len(c.getPoints())
}

func (c *cell) getPoints() []*Point {
	return c.points.points
}



type cells [][][]*cell

func (cells *cells) isFilled(x X, y Y, z Z) bool {
	return cells.at(x, y, z) != nil
}
func (cells *cells) at(x X, y Y, z Z) *cell {
	return (*cells)[int(x)][int(y)][int(z)]
}

func (cells *cells) add(point *Point) {
	cell := cells.at(point.X, point.Y, point.Z)
	cell.points.Add(point)
}

// points []*Point
func (cells *cells) init(point *Point) {
	cell := &cell{points: Points{make([]*Point, 10)}}
	(*cells)[int(point.X)][int(point.Y)][int(point.Z)] = cell
}

func (cells *cells) to(point *Point) {
	if cells.isFilled(point.X, point.Y, point.Z) {
		cells.add(point)
		return
	}
	cells.init(point)
}

func (cells *cells) atIfExists(x X, y Y, z Z) (*cell, bool) {
	if cells.isFilled(x, y, z) {
		return cells.at(x, y, z), true
	}
	return nil, false
}

func (cells *cells) pointsNum(x X, y Y, z Z) int {
	cell, isExist := cells.atIfExists(x, y, z)
	if !isExist {
		return 0
	}
	return cell.getPointNum()
}

func newCells(obje *Obje) cells {
	xsize, ysize, zsize := obje.CalcSize()
	cells := make([][][]*cell, int(xsize)+1)
	for xi := range cells {
		cells[xi] = make([][]*cell, int(ysize)+1)
		for yi := range cells[xi] {
			cells[xi][yi] = make([]*cell, int(zsize)+1)
		}
	}
	return cells
}