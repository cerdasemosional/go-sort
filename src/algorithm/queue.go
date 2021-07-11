package algorithm

type node struct {
	leftInclusive  int
	rightExclusive int
	rightNode      *node
}

type queue struct {
	leftNode  *node
	rightNode *node
}

func (state *queue) add(leftInclusive, rightExclusive int) {
	newNode := &node{leftInclusive: leftInclusive, rightExclusive: rightExclusive}

	if state.leftNode == nil {
		state.leftNode = newNode
	} else {
		state.rightNode.rightNode = newNode
	}

	state.rightNode = newNode
}

func (state *queue) addLeft(leftInclusive, rightExclusive int) {
	newNode := &node{leftInclusive: leftInclusive, rightExclusive: rightExclusive}

	if state.leftNode != nil {
		newNode.rightNode = state.leftNode
	}

	state.leftNode = newNode
}

func (state *queue) pop() *node {
	poppedNode := state.leftNode

	if poppedNode == nil {
		return poppedNode
	}

	state.leftNode = poppedNode.rightNode

	return poppedNode
}

func (state *queue) checkEmpty() bool {
	return state.leftNode == nil
}
