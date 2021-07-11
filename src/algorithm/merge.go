package algorithm

import (
	"github.com/cerdasemosional/sort-go/src/interfaces"
)

type mergeSort struct {
	points           []int32
	mergedArray      []int32
	mergedArrayIndex int
	jobs             *queue
	currentJob       *node
}

func (state *mergeSort) createMergedArray() bool {
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

func (state *mergeSort) Step() {
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

func (state *mergeSort) buildTree(leftInclusive, rightExclusive int) {
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

func mergeArray(leftArray, rightArray []int32) []int32 {
	leftArrayLength := len(leftArray)
	rightArrayLength := len(rightArray)
	combinedLength := leftArrayLength + rightArrayLength
	mergedArray := make([]int32, combinedLength)

	if leftArray[leftArrayLength-1] < rightArray[0] {
		mergedArray = append(leftArray, rightArray...)
	} else {
		leftArrayIndex := 0
		rightArrayIndex := 0
		for i := 0; i < combinedLength; i++ {
			if leftArrayIndex == leftArrayLength {
				mergedArray[i] = rightArray[rightArrayIndex]
				rightArrayIndex++
			} else if rightArrayIndex == rightArrayLength {
				mergedArray[i] = leftArray[leftArrayIndex]
				leftArrayIndex++
			} else {
				if leftArray[leftArrayIndex] < rightArray[rightArrayIndex] {
					mergedArray[i] = leftArray[leftArrayIndex]
					leftArrayIndex++
				} else {
					mergedArray[i] = rightArray[rightArrayIndex]
					rightArrayIndex++

				}
			}
		}
	}

	return mergedArray
}

func CreateMergeSort(points []int32) interfaces.AlgorithmState {
	state := mergeSort{points: points, jobs: &queue{}}
	state.buildTree(0, len(points))

	return &state
}
