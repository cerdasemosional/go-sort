package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/veandco/go-sdl2/sdl"

	"github.com/cerdasemosional/sort-go/src/algorithm"
	"github.com/cerdasemosional/sort-go/src/interfaces"
)

type config struct {
	sortAlgorithm   string
	randomAlgorithm string
	speed           int
}

func redraw(renderer *sdl.Renderer, points []int32) {
	var sdlPoints []sdl.Point
	renderer.SetDrawColor(0, 0, 0, 255)
	renderer.Clear()

	for index, point := range points {
		sdlPoints = append(sdlPoints, sdl.Point{X: int32(index), Y: int32(len(points)) - point})
	}

	renderer.SetDrawColor(255, 255, 255, 255)
	renderer.DrawPoints(sdlPoints)

	renderer.Present()
}

func generateRandom(nOfPoints int32, strategy string) []int32 {
	rand.Seed(time.Now().UnixNano())
	var randomNumbers []int32
	nOfPointsInt := int(nOfPoints)

	switch strategy {
	case "uniform":
		for i := 0; i < nOfPointsInt; i++ {
			randomNumbers = append(randomNumbers, int32(rand.Intn(nOfPointsInt)))
		}
	case "normal":
		middle := float64(nOfPoints) / 2
		for i := 0; i < nOfPointsInt; i++ {
			randomNumbers = append(randomNumbers, int32(rand.NormFloat64()*middle/3+middle))
		}
	case "cluster":
		nOfClusters := rand.Intn(10) + 5
		var clusters []int
		for i := 0; i < nOfClusters; i++ {
			clusters = append(clusters, rand.Intn(nOfPointsInt/2)+nOfPointsInt/4)
		}

		for i := 0; i < nOfClusters; i++ {
			radius := rand.Intn(nOfPointsInt/2) + 1
			for j := 0; j < nOfPointsInt/nOfClusters; j++ {
				randomNumbers = append(randomNumbers, int32(rand.Intn(radius)+clusters[i]-(radius/2)))
			}
		}

	}

	return randomNumbers
}

func getConfig() *config {
	userConfig := &config{sortAlgorithm: "insertion", randomAlgorithm: "uniform", speed: 1}

	switch length := len(os.Args); {
	case length == 2:
		userConfig.sortAlgorithm = os.Args[1]
	case length == 3:
		userConfig.sortAlgorithm = os.Args[1]
		userConfig.randomAlgorithm = os.Args[2]
	case length >= 4:
		userConfig.sortAlgorithm = os.Args[1]
		userConfig.randomAlgorithm = os.Args[2]
		speed, err := strconv.Atoi(os.Args[3])
		if err == nil {
			userConfig.speed = speed
		}
	}

	return userConfig
}

func run() int {
	const nOfPoints int32 = 1000
	winTitle := "Insertion Sort"

	window, err := sdl.CreateWindow(winTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		nOfPoints, nOfPoints, sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create window: %s\n", err)
		return 1
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create renderer: %s\n", err)
		return 2
	}
	defer renderer.Destroy()

	userConfig := getConfig()

	speed := userConfig.speed

	randoms := generateRandom(nOfPoints, userConfig.randomAlgorithm)

	var state interfaces.AlgorithmState

	switch userConfig.sortAlgorithm {
	case "insertion":
		state = algorithm.CreateInsertionSort(randoms)
	case "merge":
		state = algorithm.CreateMergeSort(randoms)
	case "quick":
		state = algorithm.CreateQuickSort(randoms)
	default:
		return 0
	}

	running := true
	prev := time.Now()
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				running = false
			}
		}

		if speed > 0 {
			for i := 0; i < speed; i++ {
				state.Step()
			}
		} else {
			now := time.Now()
			if now.Sub(prev).Milliseconds() > int64(-speed) {
				state.Step()
				prev = now
			}
		}
		redraw(renderer, randoms)
		renderer.Present()
	}

	return 0
}

func main() {
	os.Exit(run())
}
