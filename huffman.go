package main

import (
	"fmt"
	"huffman/heap"
	"huffman/node"
	"os"
)

func getFrequency(data []byte) (map[byte]int) {
    var freq = make(map[byte]int)
    for i := 0; i < len(data); i++ {
        freq[data[i]]++
    }
    return freq
}

func main() {
    var file = os.Args[1]

    compress(file);
    decompress("output.bin");
}

func decompress(filename string) {
    data, err := os.ReadFile(filename);

    if err != nil {
        fmt.Println(err)
    }

    var lenTree = int(data[0]) 
    var serializedTree = data[1:lenTree + 1]
    var tree = node.Deserialize(serializedTree)

    tree.Decompress(data[lenTree + 1:])
}


func compress(filename string) {
    data, err := os.ReadFile(filename);

    if err != nil {
        fmt.Println(err)
    } 

    var freq = getFrequency(data)

    heap := heap.Heap{};

    for char, frequency := range freq {
        heap.Insert(frequency, char)
    }

    for heap.Size() > 1 {
        heap.CombineTopTwo()
    }

    var root node.Node = heap.Pop();
    var tree = node.Huffman{Root: &root}

    var bytes = tree.Compress(data)

    var serializedTree []byte = tree.SerializeTree();
    var lenChars = []byte{byte(len(serializedTree))};

    bytes = append(serializedTree, bytes...)
    bytes = append(lenChars, bytes...) // first byte is the length of the tree

    os.WriteFile("output.bin", bytes, 0644)

    fi, err := os.Stat(filename);
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