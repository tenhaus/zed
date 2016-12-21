package main

import "errors"

type CompressedData struct {
	GridCount  uint
	Depth      uint
	Compressed map[byte]uint64
}

type Layer struct {
	Processors []Processor
}

type Processor struct {
	Points []byte
}

// Compress just gets the job done
func Compress(data []byte) error {
	var top Layer
	Partition(data, &top)

	return errors.New("nope")
}

func Partition(data []byte, layer *Layer) {
	length := len(data)
	index := 0
	points := 5

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
