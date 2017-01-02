package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

var matchCodes = []string{
	"001", "010", "100", "101", "110", "111",
}

var pointLength = 6

// CompressedData ...
type CompressedData struct {
	GridCount  uint
	Depth      uint
	Compressed map[byte]uint64
}

// Layer ...
type Layer struct {
	Processors []Processor
	Commons    Commons
}

// Row ...
type Row struct {
	Processors []Processor
}

// Processor ...
type Processor struct {
	Points []byte
	Tests  []byte
}

// Test ...
func (p *Processor) Test() string {
	var result []string

	for i := 0; i < len(p.Points); i++ {
		matchCode := Match(p.Points[i], p.Tests)
		result = append(result, matchCode)
	}

	return strings.Join(result, "")
}

// Match tests two bytes
// 000 No match
// 001 char 1
// 010 char 2
// 100 char 3
// 101 char 4
// 110 char 5
// 111 char 6
func Match(a byte, tests []byte) string {
	for i := 0; i <= 5; i++ {

		if len(tests) > i && a == tests[i] {
			return matchCodes[i]
		}
	}

	return "000"
}

// Commons is used to sort common bytes
type Commons struct {
	Data map[byte]int
	Keys []byte
}

// Slice ...
func (c Commons) Slice() [][]byte {
	var slices [][]byte

	for i := 0; i < len(c.Keys); i++ {
		if i+pointLength > len(c.Keys) {
			slices = append(slices, c.Keys[i:len(c.Keys)])
		} else {
			slices = append(slices, c.Keys[i:i+pointLength])
		}

	}

	return slices
}

func (c Commons) Len() int {
	return len(c.Keys)
}

// Swap is part of sort.Interface.
func (c Commons) Swap(i, j int) {
	c.Keys[i], c.Keys[j] = c.Keys[j], c.Keys[i]
}

func (c Commons) Less(i, j int) bool {
	return c.Data[c.Keys[i]] > c.Data[c.Keys[j]]
}

// Compress just gets the job done
func Compress(data []byte) {
	// Partition
	var top Layer
	Partition(data, &top)

	// Sort by commons
	top.Commons = MapCommons(data)
	allTests := top.Commons.Slice()

	var result []string

	for _, tests := range allTests {
		// map first layer
		for _, processor := range top.Processors {
			processor.Tests = tests
			result = append(result, processor.Test())
		}
	}

	fmt.Println(strings.Join(result, ""))
}

// MapCommons ...
func MapCommons(data []byte) Commons {

	var commons Commons
	commons.Data = make(map[byte]int)

	for _, val := range data {
		if _, ok := commons.Data[val]; !ok {
			commons.Data[val] = 1
			commons.Keys = append(commons.Keys, val)
		} else {
			commons.Data[val]++
		}
	}

	sort.Sort(commons)
	return commons
}

// GetGridSize ...
func GetGridSize(processorLength int) (int, int) {
	return pointLength * processorLength, pointLength * processorLength
}

// GetLayerSize is probably a useless function because
// we can just len(layer.Processors) after partitioning
func GetLayerSize(data []byte) uint {
	length := len(data)
	div := float64(length) / float64(pointLength)
	ceil := math.Ceil(div)
	return uint(ceil)
}

// Partition ...
func Partition(data []byte, layer *Layer) {
	length := len(data)

	for i := 0; i*pointLength < length; i++ {

		// Create the processor
		var processor Processor

		// Take a chunk
		if (i*pointLength)+pointLength > length {
			// Fill with the remaining data
			processor.Points = data[i*pointLength : length]
			// FillPartialProcessor(&processor)

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
		GenerateEmptyProcessor(&nullProcessor)
		layer.Processors = append(layer.Processors, nullProcessor)
	}
}

// GenerateEmptyProcessor ...
func GenerateEmptyProcessor(processor *Processor) {
	for i := 0; i < pointLength; i++ {
		processor.Points = append(processor.Points, 0x00)
	}
}

// FillPartialProcessor ...
func FillPartialProcessor(processor *Processor) {
	length := len(processor.Points)
	numEmpty := pointLength - length

	for i := pointLength - numEmpty; i < pointLength; i++ {
		processor.Points = append(processor.Points, 0x00)
	}
}
