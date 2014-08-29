package main

import "./vec3d"
import "./camera"
import "./models"

import "image"
import "image/png"
import "image/color"
import "os"

import "fmt"

func main() {
	cam := camera.Camera{[3]float64{20, 0, 0}, 800, 800} // camera always faces towards the origin from the positive x axis

	light := [3]float64{0, 0, 0}

	var mod []models.Model // a slice of models to be rendered
	mod = append(mod, models.Sphere{1, [3]float64{-20, 0, 0}})
	//mod = append(mod, models.Sphere{20, [3]float64{-495, 0, 0}})

	m := image.NewRGBA(image.Rect(0, 0, cam.ViewWidth, cam.ViewHeight)) // create the image to be drawn upon

	// for every pixel in the image
	for w := 0; w < cam.ViewWidth; w++ {
		for h := 0; h < cam.ViewHeight; h++ {
			viewPoint := cam.SwitchToCartesianThreeSpace(w, h)
			direction := vec3d.Sub(viewPoint, cam.Position)
			for o := 0; o < len(mod); o++ { // for all models
				tValues := mod[o].Intersect(cam.Position, direction)
				if len(tValues) == 2 {
					t := models.SmallestPositive(tValues[0], tValues[1])
					intersectPoint := vec3d.Mult(direction, t)
					lightray := vec3d.Sub(intersectPoint, light)
					normal := mod[o].Normal(intersectPoint)
					dotprod := vec3d.Dot(vec3d.Unit(lightray), vec3d.Unit(normal))
					//fmt.Println(dotprod)
					//fmt.Println("---")
					var lightWeight uint8 = 0
					if dotprod > 0 {
						lightWeight = uint8(255 * dotprod)
					}
					m.Set(w, h, color.RGBA{lightWeight, lightWeight, lightWeight, 255})
				} else {
					m.Set(w, h, color.RGBA{255, 255, 255, 255})
				}
			}
		}
	}

	defer fmt.Println("Done!")

	w, _ := os.Create("out.png")
	defer w.Close()
	png.Encode(w, m)
}