package node

import "fmt"

type Char struct {
    Char byte
    Frequency int
}

type Node struct {
    Val Char
    Left *Node
    Right *Node
}

func (n *Node) IsLeaf() (bool) {
    return n.Left == nil && n.Right == nil
}

func NewNode(val Char) (*Node) {
    var node = new(Node)
    node.Val = val
    return node
}

type Huffman struct {
    Root *Node
}

func deserializeHelper(data []byte, num *int) (*Node) {
    if(*num >= len(data)) {
        return nil
    }
    if data[*num] == 255 {
        *num = *num + 1
        return nil
    }

    var root = NewNode(Char{Char: data[*num]})
    *num = *num + 1
    root.Left = deserializeHelper(data, num)
    root.Right = deserializeHelper(data, num)
    return root
}

func Deserialize(data []byte) Huffman {
    if (len(data) == 0) {
        return Huffman{Root: nil}
    }

    var num = 0
    return Huffman{Root: deserializeHelper(data, &num) }
}

func (h *Huffman) Decompress(data []byte) {
    var lastNode = h.Root;
    
    for _, byte := range data {
        for i := 0; i < 8; i++ {
            var direction = (byte >> (7 - i)) & 1
            if (direction == 0) {
                lastNode = lastNode.Left
            } else {
                lastNode = lastNode.Right
            }
    
            if (lastNode.IsLeaf()) {
                fmt.Print(string(lastNode.Val.Char))
                lastNode = h.Root
            } 
        }
    }
}

func (h *Huffman) Compress(inputData []byte) []byte {
    var compressedData = []byte{} 
    var buffer byte;
    var count int = 0;
        
    for _, char := range inputData {
        var path = h.GetPath(char);
        var pathLen = len(path)
        for i := range path {
            var bit = path[pathLen - 1 - i]
            buffer = buffer << 1
            // set buffer position 0 to bit value
            buffer = buffer | byte(bit) 
            if count == 7 {
                compressedData = append(compressedData, buffer)
                buffer = 0;
                count = 0;
            } else {
                count++;
            }
        }
    }

    if (count != 0) {
        var shift = 8 - count;
        buffer = buffer << shift;
        compressedData = append(compressedData, buffer)
        
    }
    
    return compressedData
}

func (h *Huffman) SerializeTree() ([]byte) {
    var chars = []byte{};
    h.preorder(h.Root, &chars)

    return chars
}

func (h * Huffman) preorder(root *Node, chars *[]byte) {
    if root == nil {
        *chars = append(*chars, 255);
        return
    }

    *chars = append(*chars, root.Val.Char)
    h.preorder(root.Left, chars)
    h.preorder(root.Right, chars)
}

func (h *Huffman) GetPath(target byte) ([]int) {
    var path = []int{}
    var found bool = false;
    h.pathHelper(h.Root, target, &path, &found);
    return path
}

func (h *Huffman) pathHelper(root *Node, target byte, path *[]int, found *bool) (bool) {
    if (root == nil) {
        return false
    }

    if (root.Val.Char == target && !*found) {
        *found = true;
        return true;
    }

    var left = h.pathHelper(root.Left, target, path, found)
    var right = h.pathHelper(root.Right, target, path, found)

    if (left && *found) {
        *path = append(*path, 0)
        return true
    } 

    if (right && *found) {
        *path = append(*path, 1)
        return true
    }

    return false;
}

