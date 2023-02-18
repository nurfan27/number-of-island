package NumberOfIslands

type NumberOfIslands struct {
	Scanned [15][15]bool
	maxRow  int
	maxCol  int
}

func NewNumberOfIsland() *NumberOfIslands {
	return &NumberOfIslands{
		maxRow: 5,
		maxCol: 5,
	}
}

var (
	xAdjacentMath = []int{-1, -1, -1, 0, 0, 1, 1, 1}
	yAdjacentMath = []int{-1, 0, 1, -1, 1, -1, 0, 1}
)

func (l *NumberOfIslands) CountIslands(grid [][]int) int {
	var count = 0
	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			if grid[x][y] == 1 && !l.Scanned[x][y] {
				l.adjacentVertices(grid, x, y)
				count++
			}
		}
	}

	return count
}

func (l *NumberOfIslands) adjacentVertices(grid [][]int, x, y int) {
	// Mark cell is scanned
	l.Scanned[x][y] = true

	// 8 is adjacent vertices from cell
	for i := 0; i < 8; i++ {
		if !l.isGridLimit(x, y, xAdjacentMath[i], yAdjacentMath[i]) && grid[x][y] != 0 && !l.isScanned(x, y, i) {
			l.adjacentVertices(grid, x+xAdjacentMath[i], y+yAdjacentMath[i])
		}
	}
}

func (l *NumberOfIslands) isGridLimit(x, y, xMath, yMath int) bool {
	xLimit := x + xMath
	yLimit := y + yMath

	if xLimit >= 0 && yLimit >= 0 && x < l.maxRow && y < l.maxCol {
		return false
	}

	return true
}

func (l *NumberOfIslands) isScanned(x, y, i int) bool {
	return l.Scanned[x+xAdjacentMath[i]][y+yAdjacentMath[i]]
}
