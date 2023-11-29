# Huffman Compressor

# Description
An implementation of the Huffman Data Compressor written in Go. 

# Motivation
This data compressor was just an excuse to learn Go and practice some algorithms and data structures. The Huffman algorithm is a beautiful lossless data compression algorithm that uses a binary tree (made from a frequency analysis and a min-heap) to compress data.

# Roadmap
This is my second time implementing the Huffman Compressor. The first time, I wrote the decompression algorithm as well (after writing the compression algorithm, the decompression is trivial). For *this* project, I haven't implemented the decompression portion. 

I decided not to because decompressing a file required the tree from which it was encoded, this means that you can only decompress a file while the program is running and it still has access to the Huffman Tree. To combat this, I am working on a concise way to serialize and deserialize the binary tree for storage in the binary output file. Doing this would allow me to decompress a file even after the program exited. Although, This does mean that the effectiveness of the compression would decrease.