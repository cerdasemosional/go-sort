package algorithm

import (
	"github.com/cerdasemosional/sort-go/src/interfaces"
)

type insertionSort struct {
	points        []int32
	forwardIndex  int
	backwardIndex int
}

func (state *insertionSort) Step() {
	if state.forwardIndex <= len(state.points) {
		now := state.forwardIndex - state.backwardIndex
		prev := now - 1

		shouldEnd := true

		if state.points[now] < state.points[prev] {
			state.points[now], state.points[prev] = state.points[prev], state.points[now]
			shouldEnd = false
		}

		if prev == 0 || shouldEnd {
			state.backwardIndex = 1
			state.forwardIndex++

			return
		}
		state.backwardIndex++
	}
}

func CreateInsertionSort(points []int32) interfaces.AlgorithmState {
	return &insertionSort{points: points, forwardIndex: 1, backwardIndex: 0}
}
