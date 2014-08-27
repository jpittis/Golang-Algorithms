package main

import "./vec3d"
import "fmt"
import "math"

import "image"
import "image/png"
import "image/color"
import "os"

func intSphere(ray [3]float64) (bool, float64, float64) {
	var r float64 = 9
	var sx float64 = 0
	var sy float64 = 10
	var sz float64 = 0

	a := math.Pow(ray[0], 2) + math.Pow(ray[1], 2) + math.Pow(ray[2], 2)
	b := (2 * (ray[0] * sx)) + (2 * (ray[1] * sy)) + (2 * (ray[2] * sz))
	c := math.Pow(sx, 2) + math.Pow(sy, 2) + math.Pow(sz, 2) - math.Pow(r, 2)

	disc := math.Pow(b, 2) - (4 * a * c)

	if disc > 0 {
		t1 := (math.Sqrt(disc) - b) / (2 * a)
		t2 := (math.Sqrt(disc) + b) / (2 * a)
		return true, t1, t2
	}
	return false, 0, 0
}

func pixelPoint(r, c int) (int, int){
	r = -r
	c = c - (400)
	r = r + (400)
	return r, c
} 

func main() {
	light := [3]float64{50, 250, -100}

	cam := [3]float64{0, -100, 0}

	var width float64 = 800
	var height float64 = 800

	m := image.NewRGBA(image.Rect(0, 0, int(width), int(height)))

	for r := 0; r < 800; r++ {
		for c := 0; c < 800; c++ {

			var x, z = pixelPoint(r, c)
			view := [3]float64{float64(x), 0, float64(z)}
			ray := vec3d.Sub(view, cam)

			var intercept, t1, t2 = intSphere(ray)

			var t float64

			if t1 > 0 && t2 > 0 {
				if t1 > t2 {
					t = t2
				} else {
					t = t1
				}
			} else if t1 > 0 {
				t = t1
			} else {
				t = t2
			}

			if intercept {
				lightray := vec3d.Sub(light, vec3d.Mult(ray, t))
				sphere := [3]float64{0, 10, 0}
				normal := vec3d.Sub(ray, sphere)
				dotprod := vec3d.Dot(vec3d.Unit(lightray), vec3d.Unit(normal))
				var alpha uint8
				if dotprod > 0 {
					alpha = uint8(255 * dotprod)
				} else {
					alpha = 0
				}
				m.Set(r, c, color.RGBA{255, alpha, alpha, 255})
			}else {
				m.Set(r, c, color.RGBA{0, 70, 255, 255})
			}
		}
	}

	fmt.Println("Done!")

	w, _ := os.Create("new.png")
	defer w.Close()
	png.Encode(w, m)
}