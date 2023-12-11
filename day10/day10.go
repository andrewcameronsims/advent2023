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
	Outside
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
	Grid       [][]Tile
	Distances  map[Point]int
	HighRes    [][]Tile
}

func NewMaze(lines []string) *Maze {
	var m Maze
	m.Map = make(map[Point]Tile)
	m.Distances = make(map[Point]int)

	grid := make([][]Tile, len(lines))
	for y, line := range lines {
		row := make([]Tile, len(line))

		for x, r := range line {
			p := Point{x, y}
			t := tile(r)

			if t == Start {
				m.StartPoint = p
			}

			row[x] = t
			m.Map[p] = t
		}

		grid[y] = row
	}
	m.Grid = grid

	return &m
}

func (maze *Maze) FindInteriorTiles() int {
	// turn everything that is not main pipe to ground
	maze.contrastMainPipe()

	// extend the arrays dimensions by factor of 3
	highResGrid := make([][]Tile, 3*len(maze.Grid))
	for i := 0; i < len(highResGrid); i++ {
		highResGrid[i] = make([]Tile, 3*len(maze.Grid[0]))
	}

	for i, row := range maze.Grid {
		for j, tile := range row {
			var highResTile [][]Tile

			switch tile {
			case NS:
				highResTile = [][]Tile{
					{Ground, NS, Ground},
					{Ground, NS, Ground},
					{Ground, NS, Ground},
				}
			case WE:
				highResTile = [][]Tile{
					{Ground, Ground, Ground},
					{WE, WE, WE},
					{Ground, Ground, Ground},
				}
			case NE:
				highResTile = [][]Tile{
					{Ground, NE, Ground},
					{Ground, NE, NE},
					{Ground, Ground, Ground},
				}
			case NW:
				highResTile = [][]Tile{
					{Ground, NW, Ground},
					{NW, NW, Ground},
					{Ground, Ground, Ground},
				}
			case SW:
				highResTile = [][]Tile{
					{Ground, Ground, Ground},
					{SW, SW, Ground},
					{Ground, SW, Ground},
				}
			case SE:
				highResTile = [][]Tile{
					{Ground, Ground, Ground},
					{Ground, SE, SE},
					{Ground, SE, Ground},
				}
			case Ground:
				highResTile = [][]Tile{
					{Ground, Ground, Ground},
					{Ground, Ground, Ground},
					{Ground, Ground, Ground},
				}
			case Start:
				highResTile = [][]Tile{
					{Ground, Start, Ground},
					{Ground, Start, Start},
					{Ground, Ground, Ground},
				}
			default:
				panic(tile)
			}

			for hi := 0; hi < len(highResTile); hi++ {
				for hj := 0; hj < len(highResTile[hi]); hj++ {
					highResGrid[i*3+hi][j*3+hj] = highResTile[hi][hj]
				}
			}
		}
	}

	maze.HighRes = highResGrid

	// BFS from 0,0 to paint outside
	maze.Search(0, 0, func(i, j int) {
		if maze.HighRes[i][j] == Ground {
			maze.HighRes[i][j] = Outside
		}
	})

	var ground int
	for i := 0; i < len(maze.HighRes); i += 3 {
		for j := 0; j < len(maze.HighRes[i]); j += 3 {
			sq := maze.HighRes[i][j] + maze.HighRes[i+1][j] + maze.HighRes[i+2][j] +
				maze.HighRes[i][j+1] + maze.HighRes[i+1][j+1] + maze.HighRes[i+2][j+1] +
				maze.HighRes[i][j+2] + maze.HighRes[i+1][j+2] + maze.HighRes[i+2][j+2]

			if sq == 54 {
				ground++
			}
		}
	}

	return ground
}

func (maze *Maze) Search(startI, startJ int, op func(i, j int)) {
	visited := map[Point]bool{
		{x: startJ, y: startI}: true,
	}
	next := []Point{
		{x: 0, y: 0},
	}

	maze.search(op, visited, next)
}

func (maze *Maze) search(op func(i, j int), visited map[Point]bool, next []Point) {
	if len(next) == 0 {
		return
	}

	point := next[0]
	op(point.y, point.x)

	next = append(next[1:], maze.getHighResSurrounds(point.y, point.x, visited)...)

	maze.search(op, visited, next)
}

func (maze *Maze) getHighResSurrounds(i, j int, visited map[Point]bool) []Point {
	var ps []Point

	p := Point{x: j, y: i - 1}
	if i > 0 && !visited[p] && maze.HighRes[p.y][p.x] == Ground {
		ps = append(ps, p)
		visited[p] = true
	}

	p = Point{x: j - 1, y: i}
	if j > 0 && !visited[p] && maze.HighRes[p.y][p.x] == Ground {
		ps = append(ps, p)
		visited[p] = true
	}

	p = Point{x: j, y: i + 1}
	if i < len(maze.HighRes)-1 && !visited[p] && maze.HighRes[p.y][p.x] == Ground {
		ps = append(ps, p)
		visited[p] = true
	}

	p = Point{x: j + 1, y: i}
	if j < len(maze.HighRes[0])-1 && !visited[p] && maze.HighRes[p.y][p.x] == Ground {
		ps = append(ps, p)
		visited[p] = true
	}

	return ps
}

func (maze *Maze) contrastMainPipe() {
	for i := range maze.Grid {
		for j := range maze.Grid[i] {
			_, ok := maze.Distances[Point{x: j, y: i}]
			if !ok {
				maze.Grid[i][j] = Ground
			}
		}
	}
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
	maze.Distances[curr] = 1
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
	maze.FindDistances()

	return maze.FindInteriorTiles()
}
