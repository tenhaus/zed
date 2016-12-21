package main

import "fmt"

// CompressedData ...
type CompressedData struct {
	GridCount  uint
	Depth      uint
	Compressed map[byte]uint64
}

// Layer ...
type Layer struct {
	Processors []Processor
}

// Processor ...
type Processor struct {
	Points []byte
}

// Compress just gets the job done
func Compress(data []byte) {
	var top Layer
	Partition(data, &top)
	Spin(&top)
}

// Spin ...
func Spin(layer *Layer) {
	fmt.Println(layer)
}

// Partition ...
func Partition(data []byte, layer *Layer) {
	length := len(data)
	index := 0
	points := 6

	for {

		// Break if we're done
		if index*points >= length {
			break
		}

		// Create the processor
		var processor Processor

		// Take a chunk
		if (index*points)+points > length {
			processor.Points = data[index*points : length]
		} else {
			processor.Points = data[index*points : (index*points)+points]
		}

		// Add to the stack
		layer.Processors = append(layer.Processors, processor)

		// Continue
		index++
	}
}
