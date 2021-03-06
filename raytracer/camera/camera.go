package camera

type Camera struct {
	Position [3]float64
	ViewWidth int
	ViewHeight int
}

func (c Camera) SwitchToCartesianThreeSpace(x, y int) [3]float64 {
	var x3 float64 = 0 // view is set to the origin
	var y3 float64 = float64(x) - (float64(c.ViewWidth) / 2)
	var z3 float64 = float64(-y) + (float64(c.ViewHeight) / 2)
	return [3]float64{x3, y3, z3}
}