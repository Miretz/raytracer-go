package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const imageWidth = 256
const imageHeight = 256
const outputFile = "./output.ppm"

func writePPM() {
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

	for j := 0; j < imageHeight; j++ {
		for i := 0; i < imageWidth; i++ {
			r := float64(i) / (imageWidth - 1)
			g := float64(j) / (imageHeight - 1)
			b := 0.25
			ir := int32(255.999 * r)
			ig := int32(255.999 * g)
			ib := int32(255.999 * b)
			linesToWrite = append(linesToWrite, fmt.Sprintf("%d %d %d", ir, ig, ib))
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
