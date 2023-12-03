package day3

import (
	"fmt"
	"math"
	"unicode"
)

type Part struct {
	Symbol           rune
	SymbolI, SymbolJ int
	Number           int
}

type Schematic struct {
	Matrix [][]rune
	Parts  []Part
}

type Point struct {
	I, J int
}

func (sch *Schematic) GearRatios() int {
	var ratios int

	gears := make(map[Point][]int)
	for _, part := range sch.Parts {
		if part.Symbol == '*' {
			point := Point{part.SymbolI, part.SymbolJ}
			gears[point] = append(gears[point], part.Number)
		}
	}

	for _, ns := range gears {
		if len(ns) == 2 {
			ratio := 1

			for _, n := range ns {
				ratio *= n
			}

			ratios += ratio
		}
	}

	return ratios
}

func (sch *Schematic) FindParts() {
	var parts []Part

	m := sch.Matrix
	for i := len(m) - 1; i >= 0; i-- {
		for j := len(m[i]) - 1; j >= 0; j-- {
			var number int
			var symbol rune
			var symbolI int
			var symbolJ int

			var unit int
			for j >= 0 && unicode.IsDigit(m[i][j]) {
				if symbol == 0 {
					symbol, symbolI, symbolJ = sch.checkSurround(i, j)
				}
				value := int(m[i][j] - '0')
				number += (value * int(math.Pow(10, float64(unit))))

				j--
				unit++
			}

			if number > 0 && symbol > 0 {
				parts = append(parts, Part{Number: number, Symbol: symbol, SymbolI: symbolI, SymbolJ: symbolJ})
			}
		}
	}

	sch.Parts = parts
}

func (sch *Schematic) checkSurround(i, j int) (rune, int, int) {
	for di := -1; di < 2; di++ {
		for dj := -1; dj < 2; dj++ {
			if di == 0 && dj == 0 {
				continue
			}

			ii := i + di
			jj := j + dj

			if sch.contains(ii, jj) {
				symbol := sch.Matrix[ii][jj]
				if !unicode.IsDigit(symbol) && symbol != '.' {
					return symbol, ii, jj
				}
			}
		}
	}

	return 0, 0, 0
}

func (sch *Schematic) contains(i, j int) bool {
	return i >= 0 && j >= 0 && i < len(sch.Matrix) && j < len(sch.Matrix[i])
}

func (sch *Schematic) SumParts() int {
	var sum int

	for _, p := range sch.Parts {
		sum += p.Number
	}

	return sum
}

func Solution(lines []string) {
	partOneSolution := partOne(lines)
	fmt.Printf("Part One: %d\n", partOneSolution)

	partTwoSolution := partTwo(lines)
	fmt.Printf("Part Two: %d\n", partTwoSolution)
}

func partOne(lines []string) int {
	matrix := make([][]rune, len(lines))
	for i, line := range lines {
		matrix[i] = []rune(line)
	}

	sch := &Schematic{
		Matrix: matrix,
	}
	sch.FindParts()

	return sch.SumParts()
}

func partTwo(lines []string) int {
	matrix := make([][]rune, len(lines))
	for i, line := range lines {
		matrix[i] = []rune(line)
	}

	sch := &Schematic{
		Matrix: matrix,
	}
	sch.FindParts()

	return sch.GearRatios()
}
