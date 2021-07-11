package algorithm

import (
	"github.com/cerdasemosional/sort-go/src/interfaces"
)

type timSort struct {
	points           []int32
	mergedArray      []int32
	mergedArrayIndex int
	insertionSort    interfaces.AlgorithmState
	jobs             *queue
	currentJob       *node
}

func (state *timSort) createMergedArray() bool {
	state.currentJob = state.jobs.pop()
	length := state.currentJob.rightExclusive - state.currentJob.leftInclusive

	if length <= 1 {
		return false
	}

	middle := length/2 + state.currentJob.leftInclusive
	leftArray := make([]int32, middle-state.currentJob.leftInclusive)
	rightArray := make([]int32, state.currentJob.rightExclusive-middle)

	copy(leftArray, state.points[state.currentJob.leftInclusive:middle])
	copy(rightArray, state.points[middle:state.currentJob.rightExclusive])

	state.mergedArray = mergeArray(leftArray, rightArray)

	return true
}

func (state *timSort) Step() {
	if state.jobs.checkEmpty() && state.mergedArray == nil {
		return
	}

	if state.mergedArray == nil {
		for !state.createMergedArray() {
		}
	}

	state.points[state.mergedArrayIndex+state.currentJob.leftInclusive] = state.mergedArray[state.mergedArrayIndex]
	state.mergedArrayIndex++

	if state.mergedArrayIndex >= len(state.mergedArray) {
		state.mergedArray = nil
		state.mergedArrayIndex = 0
	}
}

func (state *timSort) buildTree(leftInclusive, rightExclusive int) {
	length := rightExclusive - leftInclusive

	if length <= 2 {
		state.jobs.add(leftInclusive, rightExclusive)
		return
	}

	middle := length/2 + leftInclusive

	state.buildTree(leftInclusive, middle)
	state.buildTree(middle, rightExclusive)

	state.jobs.add(leftInclusive, rightExclusive)
}

func CreateTimSort(points []int32) interfaces.AlgorithmState {
	state := mergeSort{points: points, jobs: &queue{}}
	state.buildTree(0, len(points))

	return &state
}
