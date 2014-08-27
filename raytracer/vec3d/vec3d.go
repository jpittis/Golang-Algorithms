package vec3d

import "math"

func Add(a, b [3]float64) [3]float64 {
	var result [3]float64
	for i := 0; i < 3; i++ {
		result[i] = a[i] + b[i]
	}
	return result
}

func Sub(a, b [3]float64) [3]float64 {
	var result [3]float64
	for i := 0; i < 3; i++ {
		result[i] = a[i] - b[i]
	}
	return result
}

func Mult(vec [3]float64, s float64) [3]float64 {
	var result [3]float64
	for i := 0; i < 3; i++ {
		result[i] = vec[i] * s
	}
	return result
}

func Unit(vec [3]float64) [3]float64 {
	return Mult(vec, 1 / Mag(vec))
}

func Mag(vec [3]float64) float64 {
	return math.Sqrt(math.Pow(vec[0], 2) + math.Pow(vec[1], 2) + math.Pow(vec[2], 2))
}

func Dot(a, b [3]float64) float64 {
	var dotProduct float64
	for i := 0; i < 3; i++ {
		dotProduct = dotProduct + (a[i] * b[i])
	}
	return dotProduct
}

func Cross(a, b [3]float64) [3]float64 {
	var result [3]float64

	result[0] = (a[1] * b[2]) - (a[2] * b[1])
	result[1] = (a[2] * b[0]) - (a[0] * b[2])
	result[2] = (a[0] * b[1]) - (a[1] * b[0])

	return result
}