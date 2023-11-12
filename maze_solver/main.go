package main

import (
	"fmt"

	"github.com/chadsmith12/go_algorithms/arraylist"
)

type Point struct {
	x int
	y int
}

var DIR = [4][2]int{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

func walk(maze []string, wall byte, current, end Point, seen [][]bool, path *arraylist.ArrayList[Point]) bool {
	// Base Case
	// we are off the map
	if current.x < 0 || current.x >= len(maze[0]) || current.y < 0 || current.y >= len(maze) {
		return false
	}

	// on a wall
	if maze[current.y][current.x] == wall {
		return false
	}

	// we have found the end
	if current.x == end.x && current.y == end.y {
		path.Push(end)
		return true
	}

	// we have already seen this point
	if seen[current.y][current.x] {
		return false
	}

	seen[current.y][current.x] = true
	path.Push(current)

	for _, col := range DIR {
		point := Point{
			x: current.x + col[0],
			y: current.y + col[1],
		}
		couldWalk := walk(maze, wall, point, end, seen, path)
		if couldWalk {
			return true
		}
	}

	path.Pop()

	return false
}

func maze_solver(maze []string, wall byte, start, end Point) *arraylist.ArrayList[Point] {
	seen := make([][]bool, len(maze))
	for i := range seen {
		seen[i] = make([]bool, len(maze[0]))
	}
	path := arraylist.NewList[Point](0)

	walk(maze, wall, start, end, seen, path)

	return path
}

func main() {
	mazeData := []string{
		"xxxxxxxxxx x",
		"x        x x",
		"x        x x",
		"x xxxxxxxx x",
		"x          x",
		"x xxxxxxxxxx",
	}

	start := Point{x: 10, y: 0}
	end := Point{x: 1, y: 5}
	path := maze_solver(mazeData, 'x', start, end)

	fmt.Println(path)
}
