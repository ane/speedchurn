package main

import (
	"runtime"
	"time"
	"sort"
)

// Churn processes a log file first by chunking it into parts and then mapreducing
// all of these simultaneously.
func Churn(file string, ch chan ChanStats) {
	specs := LineChunks(4, file)
	chunks := LoadChunks(file, specs)

	c := ChanStats{channelName: file, chunks: chunks, matcher: new(IrssiMatcher)}

	t := time.Now()
	debug.Printf("Using %d cores (%d available)\n", runtime.GOMAXPROCS(0), runtime.NumCPU())
	c.stats = MapReduce(MapChunk, ReduceChunks, GetChunks(c), 4).(StatsChunk)
	c.stats.relevant.Users = MergeSimilarNicks(c.stats.relevant, 1)
	dur := time.Since(t)
	debug.Println(file, "complete in", dur)

	ch <- c
	wg.Done()
}

// MergeSimilarNicks merges nicks that are at most minDist distance from each other.
// The distance is calculated using the Levenshtein distance of both nicks. It's usually
// smart to use minDist = 1.
func MergeSimilarNicks(r RelevantStats, minDist int) map[string]UserStats {
	nicks := []string{}
	for n, _ := range r.Users { nicks = append(nicks, n) }
	newStats := r.Users

	distances := make(map[string]map[string]int)
	toCopy := make(map[string][]string)

	for _, n1 := range nicks {
		distances[n1] = make(map[string]int)
		for _, n2 := range nicks {
			if n1 == n2 { continue; }

			dist := Levenshtein(n1, n2)
			distances[n1][n2] = dist
			if dist <= minDist {
				toCopy[n1] = append(toCopy[n1], n2)
			}
		}
	}
	nix := []string{}
	for n, _ := range toCopy { nix = append(nix, n) }
	sort.Strings(nix)

	for i := 0; i < len(nix); i++ {
		for _, n := range toCopy[nix[i]] {
			// merge
			user := newStats[nix[i]]
			user.Lines += newStats[n].Lines
			user.Words += newStats[n].Words
			newStats[nix[i]] = user
			// remove
			delete(toCopy, n)
			delete(newStats, n)
			debug.Println(nix[i], "merged with", n)
		}
	}
	return newStats
}

// Computes the smallest of each int. Used by Levenshtein.
func Min(a, b, c int) int {
	m := a
	if b < m {
		m = b
	}
	if c < m {
		m = c
	}
	return m
}

// Calculates the Levenshtein distance between a and b.
func Levenshtein(a, b string) int {
	aL := len(a)
	bL := len(b)
	if aL == 0 && bL > 0 {
		return bL;
	} else if bL == 0 && aL > 0 {
		return aL
	}

	if aL < bL { return Levenshtein(b, a) }

	d := make([]int, bL + 1)
	for i, _ := range d { d[i] = i }

	for i, xc := range a {
		x := []int{i + 1}
		for j, yc := range b {
			add := d[j + 1] + 1
			rem := x[j] + 1
			swp := d[j]
			if xc != yc {
				swp++
			}
			x = append(x, Min(add, rem, swp))
		}
		d = x
	}
	return d[len(d) - 1]
}
