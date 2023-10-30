package maze

import (
	"math/rand"

	multiarray "github.com/chadsmith12/go_algorithms/multi_array"
)

type Maze struct {
	mazeData multiarray.TwoDArray[byte]
	rows     int
	cols     int
}

type Point struct {
	x    int
	y    int
	seen bool
}

func NewPoint(x, y int) *Point {
	return &Point{x: x, y: y, seen: false}
}

func GenerateRandom(rows, cols int) *Maze {
	maze := &Maze{
		rows:     rows,
		cols:     cols,
		mazeData: *multiarray.New[byte](rows, cols),
	}

	maze.mazeData.Fill('*')
	startColumn := rand.Intn(len(maze.mazeData.Row(0)))
	endColumn := rand.Intn(len(maze.mazeData.Row(rows - 1)))
	maze.mazeData.Set(0, startColumn, ' ')
	maze.mazeData.Set(rows-1, endColumn, ' ')
	startPoint := NewPoint(0, startColumn)
	randomizeMaze(startPoint)
	return maze
}

func (m *Maze) String() string {
	return m.mazeData.String()
}

func randomizeMaze(point *Point) {
	point.seen = true

}
