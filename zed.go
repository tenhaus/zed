package main

import (
	"fmt"
	"math"
)

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
	Partition(data, &top, 6)
	Spin(&top)
}

// Spin ...
func Spin(layer *Layer) {
	fmt.Println(layer)
}

// GetLayerSize is probably a useless function because
// we can just len(layer.Processors) after partitioning
func GetLayerSize(data []byte, pointLength int) uint {
	length := len(data)
	div := float64(length) / float64(pointLength)
	ceil := math.Ceil(div)
	return uint(ceil)
}

// Partition ...
func Partition(data []byte, layer *Layer, points int) {
	length := len(data)
	index := 0

	for {

		// Break if we're done
		if index*points >= length {
			break
		}

		// Create the processor
		var processor Processor

		// Take a chunk
		if (index*points)+points > length {
			// Fill with the remaining data
			processor.Points = data[index*points : length]
			FillPartialProcessor(&processor, points)
		} else {
			processor.Points = data[index*points : (index*points)+points]
		}

		// Add to the stack
		layer.Processors = append(layer.Processors, processor)

		// Continue
		index++
	}
}

// FillPartialProcessor ...
func FillPartialProcessor(processor *Processor, pointLength int) {
	length := len(processor.Points)
	numEmpty := pointLength - length

	for i := pointLength - numEmpty; i < pointLength; i++ {
		processor.Points = append(processor.Points, 0x00)
	}
}
