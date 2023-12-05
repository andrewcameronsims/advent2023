package day5

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type Almanac struct {
	SeedToSoil            Mapping
	SoilToFertilizer      Mapping
	FertilizerToWater     Mapping
	WaterToLight          Mapping
	LightToTemperature    Mapping
	TemperatureToHumidity Mapping
	HumidityToLocation    Mapping
}

func NewAlmanac(lines []string) *Almanac {
	var almanac Almanac

	titleRe := regexp.MustCompile(`\w+-\w+-\w+`)
	mappingRe := regexp.MustCompile(`\d+`)

	for i := 0; i < len(lines); i++ {
		match := titleRe.FindString(lines[i])
		if match != "" {
			i++

			var srcs []int
			var dsts []int
			var rngs []int

			for i < len(lines) {
				matches := mappingRe.FindAllString(lines[i], -1)
				if len(matches) != 3 {
					break
				}

				dst, err := strconv.Atoi(matches[0])
				if err != nil {
					panic(err)
				}
				dsts = append(dsts, dst)

				src, err := strconv.Atoi(matches[1])
				if err != nil {
					panic(err)
				}
				srcs = append(srcs, src)

				rng, err := strconv.Atoi(matches[2])
				if err != nil {
					panic(err)
				}
				rngs = append(rngs, rng)
				i++
			}

			switch match {
			case "seed-to-soil":
				almanac.SeedToSoil = Mapping{
					SrcStart: srcs,
					DstStart: dsts,
					Range:    rngs,
				}
			case "soil-to-fertilizer":
				almanac.SoilToFertilizer = Mapping{
					SrcStart: srcs,
					DstStart: dsts,
					Range:    rngs,
				}
			case "fertilizer-to-water":
				almanac.FertilizerToWater = Mapping{
					SrcStart: srcs,
					DstStart: dsts,
					Range:    rngs,
				}
			case "water-to-light":
				almanac.WaterToLight = Mapping{
					SrcStart: srcs,
					DstStart: dsts,
					Range:    rngs,
				}
			case "light-to-temperature":
				almanac.LightToTemperature = Mapping{
					SrcStart: srcs,
					DstStart: dsts,
					Range:    rngs,
				}
			case "temperature-to-humidity":
				almanac.TemperatureToHumidity = Mapping{
					SrcStart: srcs,
					DstStart: dsts,
					Range:    rngs,
				}
			case "humidity-to-location":
				almanac.HumidityToLocation = Mapping{
					SrcStart: srcs,
					DstStart: dsts,
					Range:    rngs,
				}
			default:
				panic(match)
			}
		}
	}

	return &almanac
}

func (a *Almanac) SeedToLocation(seed int) int {
	soil := a.SeedToSoil.To(seed)
	fert := a.SoilToFertilizer.To(soil)
	water := a.FertilizerToWater.To(fert)
	light := a.WaterToLight.To(water)
	temp := a.LightToTemperature.To(light)
	hum := a.TemperatureToHumidity.To(temp)
	loc := a.HumidityToLocation.To(hum)

	return loc
}

type Mapping struct {
	SrcStart []int
	DstStart []int
	Range    []int
}

func (m *Mapping) To(src int) int {
	for i, start := range m.SrcStart {
		last := start + m.Range[i] - 1
		if src < start || src > last {
			continue
		}

		afterStart := src - start

		return m.DstStart[i] + afterStart
	}

	return src
}

func Solution(lines []string) {
	partOneSolution := partOne(lines)
	fmt.Printf("Part One: %d\n", partOneSolution)

	partTwoSolution := partTwo(lines)
	fmt.Printf("Part Two: %d\n", partTwoSolution)
}

func partTwo(lines []string) int {
	var seeds []int
	seedWords := strings.Split(lines[0], " ")[1:]

	for i := 0; i < len(seedWords); i += 2 {
		start, err := strconv.Atoi(seedWords[i])
		if err != nil {
			panic(err)
		}

		rng, err := strconv.Atoi(seedWords[i+1])
		if err != nil {
			panic(err)
		}

		for i := start; i < (start + rng - 1); i++ {
			seeds = append(seeds, i)
		}
	}

	al := NewAlmanac(lines)

	minLoc := math.MaxInt
	for _, seed := range seeds {
		loc := al.SeedToLocation(seed)
		if loc < minLoc {
			minLoc = loc
		}
	}

	return minLoc
}

func partOne(lines []string) int {
	seedWords := strings.Split(lines[0], " ")[1:]
	seeds := make([]int, len(seedWords))
	for i, sw := range seedWords {
		seed, err := strconv.Atoi(sw)
		if err != nil {
			panic(err)
		}

		seeds[i] = seed
	}

	al := NewAlmanac(lines)

	minLoc := math.MaxInt
	for _, seed := range seeds {
		loc := al.SeedToLocation(seed)
		if loc < minLoc {
			minLoc = loc
		}
	}

	return minLoc
}
