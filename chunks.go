package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type ChunkSpec struct {
	offset int64
	length int64
}

func (c ChunkSpec) String() string {
	return fmt.Sprintf("ChunkSpec { offset: %d, length: %d }", c.offset, c.length)
}

func lineChunks(numChunks int, file string) []ChunkSpec {
	fi, err := os.Open(file)
	if err != nil {
		panic("can't open file: " + file)
	}
	defer fi.Close()

	info, err := os.Stat(file)
	totalSize := info.Size()
	chunkSize := totalSize / int64(numChunks)
	//fmt.Println("File size is ", totalSize)
	//fmt.Println("Chunk size is ", chunkSize)
	specs := make([]ChunkSpec, numChunks)
	offset := int64(0)
	br := bufio.NewReader(fi)
	for i := 0; i < numChunks; i++ {
		currentChunkSize := chunkSize
		newOffset := offset + currentChunkSize
		//fmt.Println("Starting at", offset, ", seeking to ", newOffset)
		_, err := fi.Seek(newOffset, 0)
		if err == io.EOF {
			specs[i] = ChunkSpec{offset, totalSize - offset}
			break
		}

		data, err := br.ReadBytes('\n')
		untilNewLine := int64(len(data))
		currentChunkSize += untilNewLine
		//fmt.Println("Found newline at position", newOffset + untilNewLine, "(",
		//newOffset, "+", untilNewLine, ")")

		specs[i] = ChunkSpec{offset, currentChunkSize}

		offset = newOffset + untilNewLine
	}
	return specs
}
