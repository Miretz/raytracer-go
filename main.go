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

	// World
	world := hittable_list{}
	world.Add(&sphere{point3{0, 0, -1}, 0.5})
	world.Add(&sphere{point3{0, -100.5, -1}, 100})

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
				rayColor := Ray_Color(&r, &world)
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
