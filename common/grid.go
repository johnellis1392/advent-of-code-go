package common

type Grid struct {
	Width, Height int
	Start, End    Point
	matrix        [][]string
}

func createMatrix(rows, cols int) [][]string {
	matrix := make([][]string, rows)
	for r := range rows {
		matrix[r] = make([]string, cols)
	}
	return matrix
}

func NewGrid(rows, cols int) *Grid {
	return &Grid{
		Width:  cols,
		Height: rows,
		Start:  PointFromRC(0, 0),
		End:    PointFromRC(rows, cols),
		matrix: createMatrix(rows, cols),
	}
}

func NewGridFromPoints(start, end Point) *Grid {
	rows := end.R() - start.R() + 1
	cols := end.C() - start.C() + 1
	return &Grid{
		Width:  cols,
		Height: rows,
		Start:  start,
		End:    end,
		matrix: createMatrix(rows, cols),
	}
}

func (g *Grid) Size() (int, int) {
	return g.Width, g.Height
}

func (g *Grid) Contains(p Point) bool {
	return p.Within(g.Start, g.End)
}

func (g *Grid) Get(p Point) *string {
	if !g.Contains(p) {
		return nil
	}
	return &g.matrix[p.R()-g.Start.R()][p.C()-g.Start.C()]
}

func (g *Grid) Set(p Point, v string) {
	if g.Contains(p) {
		g.matrix[p.R()-g.Start.R()][p.C()-g.Start.C()] = v
	}
}
