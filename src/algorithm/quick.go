package algorithm

import (
	"github.com/cerdasemosional/sort-go/src/interfaces"
)

type quickSort struct {
	points     []int32
	jobs       *queue
	currentJob *node
	leftIndex  int
	rightIndex int
	pivot      int32
	lefTurn    bool
}

func (state *quickSort) Step() {
	if state.currentJob == nil && state.jobs.checkEmpty() {
		return
	}

	if state.currentJob == nil {
		state.currentJob = state.jobs.pop()
		length := state.currentJob.rightExclusive - state.currentJob.leftInclusive

		if length < 2 {
			state.currentJob = nil
			return
		}

		state.leftIndex = state.currentJob.leftInclusive - 1
		state.rightIndex = state.currentJob.rightExclusive
		state.lefTurn = true
		pivotIndex := (length / 2) + state.currentJob.leftInclusive
		state.pivot = state.points[pivotIndex]
	}

	if state.leftIndex >= state.rightIndex {
		state.jobs.addLeft(state.rightIndex, state.currentJob.rightExclusive)
		state.jobs.addLeft(state.currentJob.leftInclusive, state.rightIndex)
		state.currentJob = nil
		return
	}

	if state.lefTurn {
		state.leftIndex++
		if state.points[state.leftIndex] >= state.pivot {
			state.lefTurn = false
		}

		return
	}

	state.rightIndex--

	if state.points[state.rightIndex] <= state.pivot {
		state.points[state.leftIndex], state.points[state.rightIndex] = state.points[state.rightIndex], state.points[state.leftIndex]
		state.lefTurn = true
	}
}

func CreateQuickSort(points []int32) interfaces.AlgorithmState {
	state := quickSort{points: points, jobs: &queue{}}
	state.jobs.add(0, len(points))

	return &state
}
