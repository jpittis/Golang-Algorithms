package models

import ".././vec3d"

import "math"

type Model interface {
	Intersect(point, direction [3]float64) []float64
	Normal(point [3]float64) [3]float64
}

func SmallestPositive(t1, t2 float64) float64 {
	if t1 > 0 && t2 > 0 {
		if t1 > t2 {
			return t2
		} else {
			return t1
		}
	} else if t1 > 0 {
		return t1
	} else {
		return t2
	}
}

func discriminant(a, b, c float64) float64 {
	return math.Pow(b, 2) - (4 * a * c)
}

type Sphere struct {
	Radius float64
	Position [3]float64
}

func (o Sphere) Intersect(point, direction [3]float64) []float64 {
	a := math.Pow(direction[0], 2) + math.Pow(direction[1], 2) + math.Pow(direction[2], 2)
	b := (2 * (direction[0] * (o.Position[0] + point[0]))) + (2 * (direction[1] * (o.Position[1] + point[1]))) + (2 * (direction[2] * (o.Position[2] + point[2])))
	c := math.Pow(o.Position[0], 2) + math.Pow(o.Position[1], 2) + math.Pow(o.Position[2], 2) - math.Pow(o.Radius, 2)
	disc := discriminant(a, b, c)
	var result []float64
	if disc > 0 {
		result = append(result, (math.Sqrt(disc) - b) / (2 * a))
		result = append(result, (math.Sqrt(disc) + b) / (2 * a))
		return result
	}
	return result
}

func (o Sphere) Normal(point [3]float64) [3]float64 {
	return vec3d.Sub(point, o.Position)
}

/*type plane struct {
	point [3]float64
	vec1 [3]float64
	vec2 [3]float64
}

func (o plane) Intersect(point, direction [3]float64) []float64 {
	
}*/