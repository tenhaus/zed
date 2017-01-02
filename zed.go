package main

import (
	"math"
	"math/big"
	"os"
	"sort"
)

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
func (p *Processor) Test() big.Int {
	var result big.Int

	for i := 0; i < len(p.Points); i++ {
		matchCode := Match(p.Points[i], p.Tests)
		matchLen := len(matchCode)

		for j := 0; j < matchLen; j++ {
			result.SetBit(&result, (i*matchLen)+j, matchCode[j])
		}
	}

	return result
}

var noMatch = []uint{0, 0, 0}
var char1 = []uint{0, 0, 1}
var char2 = []uint{0, 1, 0}
var char3 = []uint{1, 0, 0}
var char4 = []uint{1, 0, 1}
var char5 = []uint{1, 1, 0}
var char6 = []uint{1, 1, 1}
var matchCodes = [][]uint{char1, char2, char3, char4, char5, char6}

// Match tests two bytes
func Match(a byte, tests []byte) []uint {
	for i := 0; i <= 5; i++ {
		if len(tests) > i && a == tests[i] {
			return matchCodes[i]
		}
	}

	return noMatch
}

// Commons is used to sort common bytes
type Commons struct {
	Data map[byte]int
	Keys []byte
}

// Slice ...
func (c Commons) Slice() [][]byte {
	var slices [][]byte
	numSlices := math.Ceil(float64(len(c.Keys) / pointLength))

	for i := 0; i <= int(numSlices); i++ {
		var slice []byte

		if (i*pointLength)+pointLength > len(c.Keys) {
			slice = c.Keys[(i * pointLength):len(c.Keys)]

			if i*pointLength != len(c.Keys) {
				slices = append(slices, slice)
			}

		} else {
			slice = c.Keys[(i * pointLength) : (i*pointLength)+pointLength]
			slices = append(slices, slice)
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
func Compress(data []byte, out *os.File) {
	// Partition
	var top Layer
	Partition(data, &top)

	// Sort by commons
	top.Commons = MapCommons(data)
	groupedTests := top.Commons.Slice()

	for _, tests := range groupedTests {

		// map first layer
		for _, processor := range top.Processors {
			processor.Tests = tests
			processor.Test()
			// result := processor.Test()
			// fmt.Println(result.Bits())
		}
	}
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
	numProcessors := int(math.Ceil(float64(length / pointLength)))

	for i := 0; i < numProcessors; i++ {

		// Create the processor
		var processor Processor

		// Take a chunk
		if (i*pointLength)+pointLength >= length {
			if (i * pointLength) != length {
				processor.Points = data[i*pointLength : length]
				layer.Processors = append(layer.Processors, processor)
			}
		} else {
			processor.Points = data[i*pointLength : (i*pointLength)+pointLength]
			layer.Processors = append(layer.Processors, processor)
		}
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
