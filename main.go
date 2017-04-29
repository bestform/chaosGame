package main

import (
	"math/rand"

	"time"

	"image"

	"image/color"
	"image/png"
	"os"
)

var dimX int = 1024
var dimY int = 786

type point struct {
	X int
	Y int
}

var a point
var b point
var c point

var startingPoints = map[int]*point{
	0: &a,
	1: &b,
	2: &c,
}

var current point
var img *image.RGBA

func init() {
	img = image.NewRGBA(image.Rect(0, 0, dimX, dimY))

	for x := 0; x < dimX; x++ {
		for y := 0; y < dimY; y++ {
			img.Set(x, y, color.RGBA{0, 0, 0, 255})
		}
	}

	rand.Seed(int64(time.Now().Nanosecond()))
	a = point{int(rand.Int63() % 1024), int(rand.Int63() % 768)}
	b = point{int(rand.Int63() % 1024), int(rand.Int63() % 768)}
	c = point{int(rand.Int63() % 1024), int(rand.Int63() % 768)}

	current = point{int(rand.Int63() % 1024), int(rand.Int63() % 768)}
}

func addPoint() {
	ch := rand.Int() % 3
	targetPoint := startingPoints[ch]

	newX := (current.X + targetPoint.X) / 2
	newY := (current.Y + targetPoint.Y) / 2

	current = point{newX, newY}

	img.Set(current.X, current.Y, color.RGBA{110, 110, 255, 255})
}

func main() {
	i := 0
	for {
		addPoint()
		i++
		if i == 100000 {
			break
		}
	}

	f, _ := os.OpenFile("out.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f, img)
}
