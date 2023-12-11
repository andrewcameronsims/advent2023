package day11

import (
	"fmt"
	"strings"
)

type Point struct{ i, j int }
type Boundary struct {
	Column int
	Row    int
}

type Universe struct {
	Grid                [][]rune
	Galaxies            []Point
	ExpansionBoundaries []Boundary
}

func NewUniverse(lines []string) *Universe {
	grid := make([][]rune, len(lines))
	for i, line := range lines {
		row := []rune(line)
		grid[i] = row
	}

	return &Universe{Grid: grid}
}

func (u *Universe) String() string {
	var sb strings.Builder

	for _, row := range u.Grid {
		for _, cell := range row {
			sb.WriteRune(cell)
		}
		sb.WriteRune('\n')
	}

	return sb.String()
}

func (u *Universe) MapExpansionBoundaries() {
	var expansionBoundaries []Boundary

	// Do rows first
	for i := 0; i < len(u.Grid); i++ {
		var containsGalaxy bool

		for j := range u.Grid[i] {
			if u.Grid[i][j] == '#' {
				// go to next row, this is not a candidate for expansion
				containsGalaxy = true
				break
			}
		}

		if !containsGalaxy {
			expansionBoundaries = append(expansionBoundaries, Boundary{
				Row: i,
			})
		}
	}

	// Then do columns
	for j := 0; j < len(u.Grid[0]); j++ {
		var containsGalaxy bool

		for i := 0; i < len(u.Grid); i++ {
			cell := u.Grid[i][j]
			if cell == '#' {
				containsGalaxy = true
				break
			}
		}

		if !containsGalaxy {
			expansionBoundaries = append(expansionBoundaries, Boundary{
				Column: j,
			})
		}
	}

	u.ExpansionBoundaries = expansionBoundaries
}

func (u *Universe) Expand() {
	// Do rows first
	for i := 0; i < len(u.Grid); i++ {
		var containsGalaxy bool

		for j := range u.Grid[i] {
			if u.Grid[i][j] == '#' {
				// go to next row, this is not a candidate for expansion
				containsGalaxy = true
				break
			}
		}

		if !containsGalaxy {
			// Append a new row behind this one so we don't do it twice.
			u.Grid = append(u.Grid[:i+1], u.Grid[i:]...)
			i++
		}
	}

	// Then do columns
	for j := 0; j < len(u.Grid[0]); j++ {
		var containsGalaxy bool

		for i := 0; i < len(u.Grid); i++ {
			cell := u.Grid[i][j]
			if cell == '#' {
				containsGalaxy = true
				break
			}
		}

		if !containsGalaxy {
			// insert a new cell on every row at column index j
			for i := range u.Grid {
				u.Grid[i] = append(u.Grid[i][:j+1], u.Grid[i][j:]...)
			}
			j++
		}

	}
}

func (u *Universe) MapGalaxies() {
	var gs []Point

	for i, row := range u.Grid {
		for j, cell := range row {
			if cell == '#' {
				gs = append(gs, Point{i: i, j: j})
			}
		}
	}

	u.Galaxies = gs
}

func abs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}

func max(i, j int) int {
	if i > j {
		return i
	}

	return j
}

func min(i, j int) int {
	if i < j {
		return i
	}

	return j
}

func (u *Universe) SumPairwiseLengths() int {
	var ans int
	pairsCalculated := make(map[Point]bool) // using point in a bad way to keep track

	gs := u.Galaxies

	for i := 0; i < len(gs); i++ {
		for j := 0; j < len(gs); j++ {
			if pairsCalculated[Point{i: i, j: j}] || pairsCalculated[Point{i: j, j: i}] {
				continue
			}

			// otherwise calculate manhattan distance
			g1 := gs[i]
			g2 := gs[j]

			ans += (abs(g1.i - g2.i)) + (abs(g1.j - g2.j))

			// check for each expanded boundary
			// add 999_999 for each boundary crossed
			if u.ExpansionBoundaries != nil {
				for _, b := range u.ExpansionBoundaries {
					if b.Column != 0 {
						if b.Column < max(g1.j, g2.j) && b.Column > min(g1.j, g2.j) {
							ans += 999_999
						}

						continue
					}

					if b.Row != 0 {
						if b.Row < max(g1.i, g2.i) && b.Row > min(g1.i, g2.i) {
							ans += 999_999
						}

						continue
					}
				}
			}

			pairsCalculated[Point{i: i, j: j}] = true
		}
	}

	return ans
}

func Solution(lines []string) {
	partOneSolution := partOne(lines)
	fmt.Printf("Part One: %d\n", partOneSolution)

	partTwoSolution := partTwo(lines)
	fmt.Printf("Part Two: %d\n", partTwoSolution)
}

func partOne(lines []string) int {
	u := NewUniverse(lines)
	u.Expand()
	u.MapGalaxies()

	return u.SumPairwiseLengths()
}

func partTwo(lines []string) int {
	u := NewUniverse(lines)
	u.MapGalaxies()
	u.MapExpansionBoundaries()

	return u.SumPairwiseLengths()
}
