package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func randomScene() hittable_list {
	world := hittable_list{}

	var groundMateial material = &lambertian{color{0.5, 0.5, 0.5}}
	world.Add(&sphere{point3{0, -1000, 0}, 1000, &groundMateial})

	for a := -11; a < 11; a++ {
		for b := -11; b < 11; b++ {
			chooseMat := RandomFloat()
			center := point3{float64(a) + 0.9*RandomFloat(),
				0.2, float64(b) + 0.9*RandomFloat()}

			subPoint := center.Sub(&point3{4, 0.2, 0})
			if subPoint.Length() > 0.9 {
				var sphereMaterial material = nil

				if chooseMat < 0.8 {
					// diffuse
					r1 := Vec3_Random()
					r2 := Vec3_Random()
					albedo := Vec3_Mul(&r1, &r2)
					sphereMaterial = &lambertian{albedo}
					world.Add(&sphere{center, 0.2, &sphereMaterial})
				} else if chooseMat < 0.95 {
					// metal
					albedo := Vec3_RandomBetween(0.5, 1)
					fuzz := RandomFloatBetween(0, 0.5)
					sphereMaterial = &metal{albedo, fuzz}
					world.Add(&sphere{center, 0.2, &sphereMaterial})
				} else {
					// glass
					sphereMaterial = &dielectric{1.5}
					world.Add(&sphere{center, 0.2, &sphereMaterial})
				}

			}
		}
	}

	var material1 material = &dielectric{1.5}
	world.Add(&sphere{point3{0, 1, 0}, 1.0, &material1})

	var material2 material = &lambertian{color{0.4, 0.2, 0.1}}
	world.Add(&sphere{point3{-4, 1, 0}, 1.0, &material2})

	var material3 material = &metal{color{0.7, 0.6, 0.5}, 0.0}
	world.Add(&sphere{point3{4, 1, 0}, 1.0, &material3})

	return world
}

func render() {

	// Image
	// TODO: Code needs to be optimized
	// NOTE: Currently Rendering at a lower resolution than in the book
	const aspectRatio = 3.0 / 2.0
	const imageWidth = 600 // 1200
	const imageHeight = int(imageWidth / aspectRatio)
	const outputFile = "./output.ppm"
	const samplesPerPixel = 100 // 500
	const maxDepth = 50

	// World
	world := randomScene()

	// Camera
	lookfrom := point3{13, 2, 3}
	lookat := point3{0, 0, 0}
	vup := vec3{0, 1, 0}
	distToFocus := 10.0
	aperture := 0.1
	cam := NewCamera(lookfrom, lookat, vup, 20,
		aspectRatio, aperture, distToFocus)

	// Render
	linesToWrite := make([]string, imageHeight*imageWidth)
	index := 0
	for j := imageHeight - 1; j >= 0; j-- {
		for i := 0; i < imageWidth; i++ {
			pixelColor := color{0, 0, 0}
			for s := 0; s < samplesPerPixel; s++ {
				u := (float64(i) + RandomFloat()) / float64(imageWidth-1)
				v := (float64(j) + RandomFloat()) / float64(imageHeight-1)
				r := cam.GetRay(u, v)
				rayColor := Ray_Color(&r, &world, maxDepth)
				pixelColor = pixelColor.Add(&rayColor)
			}
			linesToWrite[index] = WriteColor(&pixelColor, samplesPerPixel)
			index++
		}
	}

	// Write to file
	file, err := os.Create(outputFile)
	if err != nil {
		log.Fatal(err)
	}
	writer := bufio.NewWriter(file)
	writer.WriteString(fmt.Sprintf("P3\n%d %d\n255\n", imageWidth, imageHeight))
	for _, line := range linesToWrite {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			log.Fatalf("Error while writing to file. Err: %s", err.Error())
		}
	}
	writer.Flush()

}

func main() {
	start := time.Now()

	render()

	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println("Elapsed Time", elapsed.Seconds(), "seconds")
}
