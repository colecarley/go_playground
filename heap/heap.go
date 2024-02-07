package heap

import "huffman/node"


type Heap struct {
	nodes []node.Node
}

func (h *Heap) Size() int {
	return len(h.nodes)

}

func (h *Heap) heapify() {
	var childIndex int = len(h.nodes) - 1
    var parentIndex int = (childIndex - 1) / 2

    for parentIndex >= 0 {
        var child = h.nodes[childIndex]
        var parent = h.nodes[parentIndex]

        if child.Val.Frequency < parent.Val.Frequency {
            h.nodes[childIndex] = parent
            h.nodes[parentIndex] = child
        }

        if parentIndex == 0 {
            break
        }

        childIndex = parentIndex
        parentIndex = (childIndex - 1) / 2
    }
}

func (h *Heap) insert(node node.Node) {
	h.nodes = append(h.nodes, node)
	h.heapify()
}

func (h *Heap) Insert(Frequency int, Char byte) {
	var char = node.Char{Frequency: Frequency, Char: Char}
	var node = node.NewNode(char)
	h.insert(*node)
}

func (h *Heap) CombineTopTwo () {
	var left = h.Pop()
	var right = h.Pop()

	var char = node.Char{Frequency: left.Val.Frequency + right.Val.Frequency}
	var node = node.NewNode(char)

	node.Left = &left
	node.Right = &right

	h.insert(*node)	
}

func (h *Heap) Pop() (node.Node) {
	var node = h.nodes[0]
	h.nodes = h.nodes[1:]
	return node
}