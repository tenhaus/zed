package main

import (
	"errors"
	"fmt"
)

type CompressedData struct {
	GridCount  uint
	Depth      uint
	Compressed map[byte]uint64
}

type Processor struct {
	Points []byte
}

// Compress just gets the job done
func Compress(data []byte) error {
	Partition(data)
	return errors.New("nope")
}

func Partition(data []byte) {
	length := len(data)
	index := 0
	points := 5

	var partitions []Processor

	for {

		// Break if we're done
		if index*points >= length {
			break
		}

		// Create the processor
		var processor Processor

		if (index*points)+points > length {
			processor.Points = data[index*points : length]
			fmt.Println(index*points, length)
		} else {
			processor.Points = data[index*points : (index*points)+points]
			fmt.Println(index*points, (index*points)+points)
		}

		partitions = append(partitions, processor)
		index++
	}
}
