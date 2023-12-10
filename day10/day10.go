package day10

import "fmt"

type Point struct{ x, y int }

func (p Point) Add(o Point) Point {
	return Point{x: p.x + o.x, y: p.y + o.y}
}

type Compass int

const (
	N Compass = iota
	S
	E
	W
)

type Tile int

const (
	NS Tile = iota
	WE
	NE
	NW
	SW
	SE
	Ground
	Start
)

func tile(r rune) Tile {
	switch r {
	case '|':
		return NS
	case '-':
		return WE
	case 'L':
		return NE
	case 'J':
		return NW
	case '7':
		return SW
	case 'F':
		return SE
	case '.':
		return Ground
	case 'S':
		return Start
	default:
		panic(r)
	}
}

type Maze struct {
	StartPoint Point
	Map        map[Point]Tile
	Distances  map[Point]int
}

func NewMaze(lines []string) *Maze {
	var m Maze
	m.Map = make(map[Point]Tile)
	m.Distances = make(map[Point]int)

	for y, line := range lines {
		for x, r := range line {
			p := Point{x, y}
			t := tile(r)

			if t == Start {
				m.StartPoint = p
			}

			m.Map[p] = t
		}
	}

	return &m
}

func (maze *Maze) next(curr Point, from *Compass) Point {
	tile := maze.Map[curr]

	switch tile {
	case NS:
		if *from == N {
			return Point{x: curr.x, y: curr.y + 1}
		}

		return Point{x: curr.x, y: curr.y - 1}
	case WE:
		if *from == W {
			return Point{x: curr.x + 1, y: curr.y}
		}

		return Point{x: curr.x - 1, y: curr.y}
	case NE:
		if *from == N {
			*from = W
			return Point{x: curr.x + 1, y: curr.y}
		}

		*from = S
		return Point{x: curr.x, y: curr.y - 1}
	case NW:
		if *from == N {
			*from = E
			return Point{x: curr.x - 1, y: curr.y}
		}

		*from = S
		return Point{x: curr.x, y: curr.y - 1}
	case SW:
		if *from == S {
			*from = E
			return Point{x: curr.x - 1, y: curr.y}
		}

		*from = N
		return Point{x: curr.x, y: curr.y + 1}
	case SE:
		if *from == S {
			*from = W
			return Point{x: curr.x + 1, y: curr.y}
		}

		*from = N
		return Point{x: curr.x, y: curr.y + 1}
	default:
		panic(curr)
	}
}

func (maze *Maze) FindDistances() {
	start := maze.StartPoint
	entries := []Point{
		start.Add(Point{x: 0, y: -1}),
		start.Add(Point{x: 1, y: 0}),
	}

	maze.Distances[start] = 0

	curr := entries[0]
	from := S
	steps := 1
	for curr != start {
		curr = maze.next(curr, &from)
		maze.Distances[curr] = steps
		steps++
	}
}

func Solution(lines []string) {
	partOneSolution := partOne(lines)
	fmt.Printf("Part One: %d\n", partOneSolution)

	partTwoSolution := partTwo(lines)
	fmt.Printf("Part Two: %d\n", partTwoSolution)
}

func partOne(lines []string) int {
	maze := NewMaze(lines)
	maze.FindDistances()

	var max int
	for _, d := range maze.Distances {
		if d > max {
			max = d
		}
	}

	return max/2 + 1
}

func partTwo(lines []string) int {
	maze := NewMaze(lines)
}
