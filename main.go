package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func writePPM() {

	// Image
	const aspectRatio = 16.0 / 9.0
	const imageWidth = 400
	const imageHeight = int(imageWidth / aspectRatio)
	const outputFile = "./output.ppm"
	const samplesPerPixel = 100
	const maxDepth = 50

	// World
	world := hittable_list{}
	var materialGround material = &lambertian{color{0.8, 0.8, 0.0}}
	var materialCenter material = &lambertian{color{0.7, 0.3, 0.3}}
	var materialLeft material = &metal{color{0.8, 0.8, 0.8}, 0.3}
	var materialRight material = &metal{color{0.8, 0.6, 0.2}, 1.0}

	world.Add(&sphere{point3{0.0, -100.5, -1.0}, 100.0, &materialGround})
	world.Add(&sphere{point3{0.0, 0.0, -1.0}, 0.5, &materialCenter})
	world.Add(&sphere{point3{-1.0, 0.0, -1.0}, 0.5, &materialLeft})
	world.Add(&sphere{point3{1.0, 0.0, -1.0}, 0.5, &materialRight})

	// Camera
	cam := NewCamera()

	// Render
	file, err := os.Create(outputFile)
	if err != nil {
		log.Fatal(err)
	}
	writer := bufio.NewWriter(file)

	linesToWrite := []string{
		"P3",
		fmt.Sprintf("%d %d", imageWidth, imageHeight),
		"255",
	}

	for j := imageHeight; j >= 0; j-- {
		for i := 0; i < imageWidth; i++ {
			pixelColor := color{0, 0, 0}
			for s := 0; s < samplesPerPixel; s++ {
				u := (float64(i) + RandomFloat()) / float64(imageWidth-1)
				v := (float64(j) + RandomFloat()) / float64(imageHeight-1)
				r := cam.GetRay(u, v)
				rayColor := Ray_Color(&r, &world, maxDepth)
				pixelColor = pixelColor.Add(&rayColor)
			}
			toWrite := WriteColor(&pixelColor, samplesPerPixel)
			linesToWrite = append(linesToWrite, toWrite)
		}
	}

	for _, line := range linesToWrite {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			log.Fatalf("Error while writing to file. Err: %s", err.Error())
		}
	}
	writer.Flush()
}

func main() {
	writePPM()
}
