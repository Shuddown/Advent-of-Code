package main

import (
	"bufio"
	"cmp"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/Shuddown/Advent-of-Code/utils"
)

type point struct {
	x int
	y int
	z int
}

type graph struct {
	points []*point
}

var points []*point

func dist(a, b *point) float64 {
	dx := float64(a.x - b.x)
	dy := float64(a.y - b.y)
	dz := float64(a.z - b.z)
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

func getClosest(start, end int) float64 {
	if end <= start {
		return math.MaxFloat64
	}
	mid := (start + end) / 2
	dividingX := (points[mid].x + points[mid+1].x) / 2
	d1 := getClosest(start, mid)
	d2 := getClosest(mid+1, end)
	minD := min(d1, d2)
	mergePoints := make([]*point, 0)
	for i := mid; i > -1; i-- {
		if float64(dividingX-points[i].x) < minD {
			mergePoints = append(mergePoints, points[i])
			continue
		}
		break
	}
	for i := mid + 1; i < len(points); i++ {
		if float64(-dividingX+points[i].x) < minD {
			mergePoints = append(mergePoints, points[i])
			continue
		}
		break
	}
	slices.SortFunc(mergePoints, func(a, b *point) int {
		return cmp.Compare(a.y, b.y)
	})
	for i := range mergePoints {
		for j := i - 1; j > -1; j-- {
			if float64(mergePoints[i].y-mergePoints[j].y) < minD {
				d := dist(mergePoints[i], mergePoints[j])
				minD = min(minD, d)
				continue
			}
			break
		}
	}
	return minD
}

func main() {
	inputFile, err := utils.GetInput(8)
	utils.HandleError(err)
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		pointString := scanner.Text()
		pointStringSlice := strings.Split(pointString, ",")
		x, _ := strconv.Atoi(pointStringSlice[0])
		y, _ := strconv.Atoi(pointStringSlice[1])
		z, _ := strconv.Atoi(pointStringSlice[2])
		points = append(points, &point{x, y, z})
	}
	utils.HandleError(scanner.Err())
	slices.SortFunc(points, func(a, b *point) int {
		return cmp.Compare(a.x, b.x)
	})
	d := getClosest(0, len(points)-1)
	fmt.Println(d)
}
