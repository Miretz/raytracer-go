package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"sync"
	"time"
)

// Image
const aspectRatio = 3.0 / 2.0
const imageWidth = 600
const imageHeight = int(imageWidth / aspectRatio)
const outputFile = "./output.ppm"
const samplesPerPixel = 100
const maxDepth = 50

func randomScene() hittable_list {
	world := hittable_list{}

	var groundMateial material = &lambertian{color{0.5, 0.5, 0.5}}
	world.Add(&sphere{point3{0, -1000, 0}, 1000, &groundMateial})

	maxDist := point3{4, 0.2, 0}

	for a := -11; a < 11; a++ {
		for b := -11; b < 11; b++ {
			chooseMat := rand.Float64()
			center := point3{float64(a) + 0.9*rand.Float64(),
				0.2, float64(b) + 0.9*rand.Float64()}

			subPoint := Vec3_Sub(&center, &maxDist)
			if subPoint.Length() > 0.9 {
				var sphereMaterial material = nil

				if chooseMat < 0.8 {
					// diffuse
					r1 := Vec3_Random()
					r2 := Vec3_Random()
					albedo := Vec3_Mul(r1, r2)
					sphereMaterial = &lambertian{*albedo}
					world.Add(&sphere{center, 0.2, &sphereMaterial})
				} else if chooseMat < 0.95 {
					// metal
					albedo := Vec3_RandomBetween(0.5, 1)
					fuzz := RandomFloatBetween(0, 0.5)
					sphereMaterial = &metal{*albedo, fuzz}
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

type renderPixelResult struct {
	id    int
	value string
}

func renderPixel(
	wg *sync.WaitGroup,
	j int,
	world *hittable_list,
	cam *camera,
	result chan<- renderPixelResult) {
	defer wg.Done()
	const divU = float64(imageWidth - 1)
	const divV = float64(imageHeight - 1)
	builder := strings.Builder{}
	for i := 0; i < imageWidth; i++ {
		pixelColor := color{0, 0, 0}
		for s := 0; s < samplesPerPixel; s++ {
			u := (float64(i) + rand.Float64()) / divU
			v := (float64(j) + rand.Float64()) / divV
			pixelColor.AddAssign(Ray_Color(cam.GetRay(u, v), world, maxDepth))
		}
		builder.WriteString(WriteColor(&pixelColor, samplesPerPixel))
	}
	result <- renderPixelResult{imageHeight - j - 1, builder.String()}
}

func Render() {

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
	var wg sync.WaitGroup
	resultChan := make(chan renderPixelResult, imageHeight)
	for j := imageHeight - 1; j >= 0; j-- {
		wg.Add(1)
		go renderPixel(&wg, j, &world, &cam, resultChan)
	}
	wg.Wait()

	// Collect the results
	linesToWrite := make([]string, imageHeight)
	for i := 0; i < imageHeight; i++ {
		res := <-resultChan
		linesToWrite[res.id] = res.value
	}

	// Write to file
	file, err := os.Create(outputFile)
	if err != nil {
		log.Fatal(err)
	}
	writer := bufio.NewWriter(file)
	writer.WriteString(fmt.Sprintf("P3\n%d %d\n255\n", imageWidth, imageHeight))
	for _, line := range linesToWrite {
		_, err := writer.WriteString(line)
		if err != nil {
			log.Fatalf("Error while writing to file. Err: %s", err.Error())
		}
	}
	writer.Flush()
}

func main() {
	start := time.Now()

	Render()

	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println("Elapsed Time", elapsed.Seconds(), "seconds")
}
