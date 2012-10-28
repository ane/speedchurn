package main

import (
	"os"
	"fmt"
	"io"
	"bufio"
)

func processLog(file string, c chan int) {
	fi, err := os.Open(file)
	if err != nil { panic("can't open file:" + file) }
	defer fi.Close()

	r := bufio.NewReader(fi)
	buf := make([]byte, 1024)
	fmt.Printf("Processing... ")

	for {
		n, err := r.Read(buf)
		if err != nil && err != io.EOF { panic(err) }

		if n == 0 { break }

		// do something
		fmt.Printf(string(buf[:n]));
	}

	fmt.Printf("Done.\n")
	c <- 1
}

func main() {
	args := os.Args

	if len(args) != 2 { panic("Usage: risa <log>") }

	cs := lineChunks(4, args[1])
	for _, spec := range cs {
		fmt.Println(spec)
	}
}
