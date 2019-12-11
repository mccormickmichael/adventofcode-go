package day10

import (
	"fmt"
	"math"
	"sort"
)

type polar struct {
	asteroid
	r     float64
	theta float64
	thetaRep string
}

type radial []polar

func (r radial) Len() int {
	return len(r)
}
func (r radial) Less(i, j int) bool {
	return r[i].r < r[j].r
}

func (r radial) Swap(i, j int) {
	t := r[i]
	r[i] = r[j]
	r[j] = t
}

type targets []radial

func (t targets) Len() int {
	return len(t)
}
func (t targets) Less(i, j int) bool {
	return t[i][0].theta < t[j][0].theta
}
func (t targets) Swap(i, j int) {
	tmp := t[i]
	t[i] = t[j]
	t[j] = tmp
}

func toPolar(origin asteroid, a asteroid) polar {
	r := math.Sqrt(float64((a.x-origin.x)*(a.x-origin.x) + (a.y-origin.y)*(a.y-origin.y)))
	// X and Y are reversed in the call to Atan2 because we are sweeping clockwise from vertical
	theta := 180.0 / math.Pi * math.Atan2(float64(a.x-origin.x), -float64(a.y-origin.y))
	if theta < 0 {
		theta += 360.0
	}
	return polar{a,r, theta, fmt.Sprintf("%07.3f", theta)}
}

func collect(polars []polar) targets {
	radialMap := make(map[string]radial)
	for _, p := range polars {
		r := radialMap[p.thetaRep]
		radialMap[p.thetaRep] = append(r, p)
	}
	targets := make(targets, 0)
	for _, p := range radialMap {
		sort.Sort(p)
		targets = append(targets, p)
	}
	sort.Sort(targets)
	return targets
}