package main

import (
	"io/ioutil"
	"testing"
)

func TestCompress(t *testing.T) {
	file, _ := ioutil.ReadFile("files/simple.txt")
	Compress(file)
}

func TestFill(t *testing.T) {
	pointLength := 6
	var processor Processor

	processor.Points = append(processor.Points, 0x01)
	processor.Points = append(processor.Points, 0x01)
	processor.Points = append(processor.Points, 0x01)

	FillPartialProcessor(&processor, pointLength)

	if processor.Points[3] != 0x00 &&
		processor.Points[4] != 0x00 &&
		processor.Points[5] != 0x00 {
		t.Fail()
	}

	if processor.Points[0] != 0x01 &&
		processor.Points[1] != 0x01 &&
		processor.Points[2] != 0x01 {
		t.Fail()
	}
}

func TestPartitionsAreEven(t *testing.T) {
	file, _ := ioutil.ReadFile("files/5partitions.txt")
	var layer Layer
	Partition(file, &layer, 6)

	if len(layer.Processors) != 6 {
		t.Fail()
	}
}

func TestGetLayerSize(t *testing.T) {
	file, _ := ioutil.ReadFile("files/20chars.txt")
	layerSize := GetLayerSize(file, 6)

	if layerSize != 4 {
		t.Fail()
	}
}
