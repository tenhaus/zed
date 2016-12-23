package main

import "math"

// CompressedData ...
type CompressedData struct {
	GridCount  uint
	Depth      uint
	Compressed map[byte]uint64
}

// Layer ...
type Layer struct {
	Processors  []Processor
	PointLength int
}

// Row ...
type Row struct {
	Processors []Processor
}

// Processor ...
type Processor struct {
	Points []byte
	Angle  uint
}

// Compress just gets the job done
func Compress(data []byte) {
	var top Layer
	Partition(data, &top, 6)
	Spin(&top)
}

// Spin ...
func Spin(layer *Layer) {
	// fmt.Println(layer)
}

// GetGridSize ...
func GetGridSize(pointLength int, processorLength int) (int, int) {
	return pointLength * processorLength, pointLength * processorLength
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
func Partition(data []byte, layer *Layer, pointLength int) {
	layer.PointLength = pointLength
	length := len(data)

	for i := 0; i*pointLength < length; i++ {

		// Create the processor
		var processor Processor

		// Take a chunk
		if (i*pointLength)+pointLength > length {
			// Fill with the remaining data
			processor.Points = data[i*pointLength : length]
			FillPartialProcessor(&processor, pointLength)

		} else {
			processor.Points = data[i*pointLength : (i*pointLength)+pointLength]
		}

		// Add to the stack
		layer.Processors = append(layer.Processors, processor)
	}

	numProcessors := len(layer.Processors)
	remainder := math.Mod(float64(numProcessors), 2.0)

	// Make sure we have an odd number of processors
	if remainder == 0 {
		var nullProcessor Processor
		GenerateEmptyProcessor(&nullProcessor, pointLength)
		layer.Processors = append(layer.Processors, nullProcessor)
	}
}

// GenerateEmptyProcessor ...
func GenerateEmptyProcessor(processor *Processor, pointLength int) {
	for i := 0; i < pointLength; i++ {
		processor.Points = append(processor.Points, 0x00)
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
