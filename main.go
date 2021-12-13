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
	const imageWidth = 1920
	const imageHeight = int(imageWidth / aspectRatio)
	const outputFile = "./output.ppm"

	// World
	world := hittable_list{}
	world.Add(&sphere{point3{0, 0, -1}, 0.5})
	world.Add(&sphere{point3{0, -100.5, -1}, 100})

	// Camera
	viewportHeight := 2.0
	viewportWidth := aspectRatio * viewportHeight
	focalLength := 1.0

	origin := point3{0, 0, 0}
	horizontal := vec3{viewportWidth, 0, 0}
	vertical := vec3{0, viewportHeight, 0}

	lowerLeftCorner := origin.SubMultiple(
		horizontal.Div(2.0),
		vertical.Div(2.0),
		vec3{0, 0, focalLength})

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
			u := float64(i) / float64(imageWidth-1)
			v := float64(j) / float64(imageHeight-1)
			r := ray{origin, Vec3_AddMultiple(
				lowerLeftCorner,
				horizontal.Mul(u),
				vertical.Mul(v),
				origin.Neg())}
			pixelColor := Ray_Color(&r, &world)
			linesToWrite = append(linesToWrite, WriteColor(&pixelColor))
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
