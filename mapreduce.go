// Copyright 2012 Antoine Kalmbach <ane@iki.fi>
// Use of this source code is governed by a GPLv2 license
// found in the LICENSE file.
//
// adapted from https://github.com/dbravender/go_mapreduce/blob/master/src/mapreduce.go

package main

func MapReduce(mapper func(interface{}, chan interface{}),
	reducer func(chan interface{}, chan interface{}),
	input chan interface{},
	pool_size int) interface{} {
	reduce_input := make(chan interface{})
	reduce_output := make(chan interface{})
	worker_output := make(chan chan interface{}, pool_size)
	go reducer(reduce_input, reduce_output)
	go func() {
		for worker_chan := range worker_output {
			reduce_input <- <-worker_chan
		}
		close(reduce_input)
	}()
	go func() {
		for item := range input {
			my_chan := make(chan interface{})
			go mapper(item, my_chan)
			worker_output <- my_chan
		}
		close(worker_output)
	}()
	return <-reduce_output
}
