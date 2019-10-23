package hrtree

import "math"

// Rect represents a 2D Rectangle
type Rect struct {
	Xmin float64
	Xmax float64
	Ymin float64
	Ymax float64
}

func (r1 Rect) area() float64 {
	return (r1.Xmax - r1.Xmin) * (r1.Ymax - r1.Ymin)
}

func (r1 Rect) intersect(r2 Rect) bool {
	return r1.Xmin < r2.Xmax && r1.Xmax > r2.Xmin && r1.Ymin < r2.Ymax && r1.Ymax > r2.Ymin
}
func (r1 Rect) union(r2 Rect) Rect {
	return Rect{
		math.Min(r1.Xmin, r2.Xmin),
		math.Max(r1.Xmax, r2.Xmax),
		math.Min(r1.Ymin, r2.Ymin),
		math.Max(r1.Ymax, r2.Ymax),
	}
}
