// Copyright 2012 Antoine Kalmbach <ane@iki.fi>
// Use of this source code is governed by a GPLv2 license
// found in the LICENSE file.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type Chunk struct {
	Data  []byte
	Order int
}

type ChunkSpec struct {
	offset int64
	length int64
}

func (c ChunkSpec) String() string {
	return fmt.Sprintf("ChunkSpec { offset: %d, length: %d }", c.offset, c.length)
}

// LineChunks splits the file into numChunks parts, that begin on a newline.
// The chunks are guaranteed to end on a newline for all but the last chunk.
func LineChunks(numChunks int, file string) []ChunkSpec {
	fi, err := os.Open(file)
	if err != nil {
		panic("can't open file: " + file)
	}
	defer fi.Close()

	info, err := os.Stat(file)
	totalSize := info.Size()
	chunkSize := totalSize / int64(numChunks)
	specs := make([]ChunkSpec, numChunks)
	offset := int64(0)
	br := bufio.NewReader(fi)

	for i := 0; i < numChunks; i++ {
		currentChunkSize := chunkSize
		newOffset := offset + currentChunkSize

		_, err := fi.Seek(newOffset, 0)
		if err == io.EOF {
			specs[i] = ChunkSpec{offset, totalSize - offset}
			break
		}

		// find newline
		data, err := br.ReadBytes('\n')
		untilNewLine := int64(len(data))
		currentChunkSize += untilNewLine

		specs[i] = ChunkSpec{offset, currentChunkSize}

		offset = newOffset + untilNewLine
	}
	return specs
}

func LoadChunks(file string, specs []ChunkSpec) []Chunk {
	chunks := make([]Chunk, len(specs))
	fi, err := os.Open(file)
	if err != nil {
		panic("can't open file: " + file)
	}
	defer fi.Close()

	reader := bufio.NewReader(fi)

	for i, spec := range specs {
		chunkbuf := make([]byte, spec.length)
		fi.Seek(spec.offset, 0)
		_, err := reader.Read(chunkbuf)
		if err != nil {
			panic("error reading chunk " + string(i))
		}

		chunks[i] = Chunk{Data: chunkbuf, Order: i}
	}
	return chunks
}
