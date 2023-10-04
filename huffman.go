package main

import (
	"fmt"
	"os"
)

type Char struct {
    char byte
    Frequency int
}

type Node struct {
    Val Char
    Left *Node
    Right *Node
}

func newNode(val Char) (*Node) {
    var node = new(Node)
    node.Val = val
    return node
}

func addNode(root *Node, val Char) (*Node) {
    if root == nil {
        var newNode = new(Node)
        newNode.Val = val
        return newNode;
    }

    if val.Frequency < root.Val.Frequency {
        root.Left = addNode(root.Left, val)
        return root;
    } else {
        root.Right = addNode(root.Right, val)
        return root;
    }
}

func getPath(root *Node, target byte) ([]int) {
    var path = []int{}
    var found bool = false;
    pathHelper(root, target, &path, &found);
    return path
}

func pathHelper(root *Node, target byte, path *[]int, found *bool) (bool) {
    if (root == nil) {
        return false
    }

    if (root.Val.char == target) {
        *found = true;
        return true;
    }

    var left = pathHelper(root.Left, target, path, found)
    var right = pathHelper(root.Right, target, path, found)

    if (left) {
        *path = append(*path, 0)
        return true
    } 
    if (right) {
        *path = append(*path, 1)
        return true
    }

    return false;
}

func getFrequency(data []byte) (map[byte]int) {
    var freq = make(map[byte]int)
    for i := 0; i < len(data); i++ {
        freq[data[i]]++
    }
    return freq
}

func heapify(heap *[]Node) {
    var childIndex int = len(*heap) - 1
    var parentIndex int = (childIndex - 1) / 2

    for parentIndex >= 0 {
        var child = (*heap)[childIndex]
        var parent = (*heap)[parentIndex]

        if child.Val.Frequency < parent.Val.Frequency {
            (*heap)[childIndex] = parent
            (*heap)[parentIndex] = child
        }

        if(parentIndex == 0) {
            break
        }

        childIndex = parentIndex
        parentIndex = (childIndex - 1) / 2
    }
}


func main() {
    var file = os.Args[1]
    compress(file);
}


func compress(file string) {
    data, err := os.ReadFile(file);

    if err != nil {
        fmt.Println(err)
    } 

    var freq = getFrequency(data)

    var heap = []Node{}

    for key, value := range freq {
        var char = Char{Frequency: value, char: key}
        var node = newNode(char)
        heap = append(heap, *node)
        heapify(&heap)
    }

    for len(heap) > 1 {
        var left = heap[0]
        heap = heap[1:]
        var right = heap[0]
        heap = heap[1:]

        var char = Char{Frequency: left.Val.Frequency + right.Val.Frequency}
        var node = newNode(char)

        node.Left = &left
        node.Right = &right

        heap = append(heap, *node)
    }

    var tree = heap[0];

    var bytes = []byte{}
    var buffer byte;
    var count int = 0;
    for i := 0; i < len(data); i++ {
        var path = getPath(&tree, data[i]);
        for j := 0; j < len(path); j++ {
            buffer = buffer << 1
            buffer = buffer | byte(path[j])
            if count == 7 {
                bytes = append(bytes, buffer)
                buffer = 0
                count = 0
            } else {
                count++
            }
        }
    }


    os.WriteFile("output.bin", bytes, 0644)

    fi, err := os.Stat(file);
    if err != nil {
        return;
    }
    fiOutput, errOutput := os.Stat("output.bin");
    if errOutput != nil {
        return;
    }

    fmt.Println(fi.Size());
    fmt.Println(fiOutput.Size()); 
}