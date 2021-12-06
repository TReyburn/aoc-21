package day5

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Point struct {
	x   int
	y   int
	str string
}

func NewPoint(p string) (*Point, error) {
	xy := strings.Split(p, ",")
	if len(xy) != 2 {
		return nil, errors.New(fmt.Sprintf("Error: %v - > %v", p, xy))
	}
	x, err := strconv.Atoi(xy[0])
	if err != nil {
		return nil, err
	}
	y, err := strconv.Atoi(xy[1])
	if err != nil {
		return nil, err
	}
	return &Point{
		x:   x,
		y:   y,
		str: p,
	}, nil
}

type Grid struct {
	points map[string]*Point
	counts map[*Point]int
}

func NewGrid() Grid {
	p := make(map[string]*Point)
	c := make(map[*Point]int)

	return Grid{points: p, counts: c}
}

func (g *Grid) AddPoint(p *Point) {
	point, ok := g.points[p.str]
	if !ok {
		g.points[p.str] = p
		g.counts[p]++
	} else {
		g.counts[point]++
	}
}

//AddLine expects an input string structured as `x1,y1 -> x2,y2` where x and y are int coordinates for a point
func (g *Grid) AddLine(line string) error {
	split := strings.Split(line, " -> ")
	startPoint, err := NewPoint(split[0])
	if err != nil {
		return err
	}

	endPoint, err := NewPoint(split[1])
	if err != nil {
		return err
	}

	x1 := startPoint.x
	y1 := startPoint.y

	x2 := endPoint.x
	y2 := endPoint.y

	// this is where we exclude diagonals
	if x1 == x2 || y1 == y2 {
		g.AddPoint(startPoint)
		g.AddPoint(endPoint)
		switch {
		case x1 != x2:
			minX := min(x1, x2)
			maxX := max(x1, x2)
			diff := maxX - minX
			for i := 1; i < diff; i++ {
				p, err := NewPoint(fmt.Sprintf("%v,%v", i+minX, y1))
				if err != nil {
					return err
				}
				g.AddPoint(p)
			}
		case y1 != y2:
			minY := min(y1, y2)
			maxY := max(y1, y2)
			diff := maxY - minY
			for i := 1; i < diff; i++ {
				p, err := NewPoint(fmt.Sprintf("%v,%v", x1, i+minY))
				if err != nil {
					return err
				}
				g.AddPoint(p)
			}
		default:
			return errors.New(fmt.Sprintf("line appears to be a point: %v, %v", startPoint, endPoint))
		}
	}
	return nil
}

// BulkAddLines expects an identical string structure to AddLine - however this accepts a series of these strings separated by newline
func (g *Grid) BulkAddLines(l string) error {
	split := strings.Split(l, "\n")
	for _, line := range split {
		if line != "" {
			err := g.AddLine(line)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (g *Grid) CountOverlaps() int {
	var overlap int
	for _, v := range g.counts {
		if v > 1 {
			overlap++
		}
	}
	return overlap
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
