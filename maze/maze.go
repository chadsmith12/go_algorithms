package maze

import (
	"math/rand"

	multiarray "github.com/chadsmith12/go_algorithms/multi_array"
)

var dir = [4][2]int{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

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
	seen := multiarray.New[bool](rows, cols)
	seen.Fill(false)

	maze.mazeData.Fill('*')
	startColumn := rand.Intn(len(maze.mazeData.Row(0)))
	endColumn := rand.Intn(len(maze.mazeData.Row(rows - 1)))
	maze.mazeData.Set(0, startColumn, ' ')
	maze.mazeData.Set(rows-1, endColumn, ' ')
	startPoint := NewPoint(0, startColumn)
	maze.randomizeMaze(startPoint, seen)
	return maze
}

func (m *Maze) String() string {
	return m.mazeData.String()
}

func (m *Maze) randomizeMaze(point *Point, seen *multiarray.TwoDArray[bool]) {
	seen.Set(point.y, point.x, true)

	for {
		nextPoint, ok := randomUnvisitedNeighbor(point, seen, m)
		if !ok {
			return
		}
		m.mazeData.Set(nextPoint.y, nextPoint.x, ' ')
		m.randomizeMaze(&nextPoint, seen)
	}
}

func randomUnvisitedNeighbor(point *Point, seen *multiarray.TwoDArray[bool], maze *Maze) (Point, bool) {
	neighbors := [4]Point{}
	for i, col := range dir {
		next := Point{x: point.x + col[0], y: point.y + col[1]}
		neighbors[i] = next
	}

	for _, neighbor := range neighbors {
		if neighbor.x < 0 || neighbor.x > maze.cols-1 || neighbor.y < 0 || neighbor.y > maze.rows-1 {
			continue
		}

		if seen.Get(neighbor.y, neighbor.x) == false {
			return neighbor, true
		}
	}

	return Point{}, false
}
