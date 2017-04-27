package main

import (
	"math/rand"

	"math"

	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

var dimX int = 1024
var dimY int = 786
var a pixel.Vec
var b pixel.Vec
var c pixel.Vec
var startingPoints = map[int]*pixel.Vec{
	0: &a,
	1: &b,
	2: &c,
}
var r float64 = 1
var t float64 = 6
var current pixel.Vec

var imd *imdraw.IMDraw

func init() {
	imd = imdraw.New(nil)

	imd.Color(pixel.RGB(0.2, 0.2, 0.2))
	imd.Push(a)
	imd.Push(b)
	imd.Push(c)
	imd.Circle(r, t)
	imd.Color(pixel.RGB(0.1, 0.1, 0.9))

	rand.Seed(int64(time.Now().Nanosecond()))
	a = pixel.V(float64(rand.Int63()%1024), float64(rand.Int63()%768))
	b = pixel.V(float64(rand.Int63()%1024), float64(rand.Int63()%768))
	c = pixel.V(float64(rand.Int63()%1024), float64(rand.Int63()%768))
	current = pixel.V(float64(rand.Int63()%1024), float64(rand.Int63()%768))
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Chaos Game",
		Bounds: pixel.R(0, 0, float64(dimX), float64(dimY)),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	for !win.Closed() {
		win.Clear(colornames.Black)
		imd.Draw(win)
		win.Update()
	}
}

func addPoint(imd *imdraw.IMDraw) {
	ch := rand.Int() % 3
	targetPoint := startingPoints[ch]

	newX := math.Abs(current.X()+targetPoint.X()) / 2
	newY := math.Abs(current.Y()+targetPoint.Y()) / 2

	current = pixel.V(newX, newY)

	imd.Push(current)
	imd.Circle(r, t/4)
}

func main() {
	i := 0
	for {
		addPoint(imd)
		i++
		if i == 5000 {
			break
		}
	}
	pixelgl.Run(run)
}
