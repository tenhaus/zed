package main

import (
	"fmt"
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

func TestPartitionsAreOdd(t *testing.T) {
	file, _ := ioutil.ReadFile("files/20chars.txt")
	var layer Layer
	Partition(file, &layer, 6)

	if len(layer.Processors) != 5 {
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

func TestGetGridSize(t *testing.T) {
	file, _ := ioutil.ReadFile("files/5processors.txt")
	var layer Layer
	Partition(file, &layer, 6)
	xMax, yMax := GetGridSize(6, len(layer.Processors))

	if xMax != 30 && yMax != 30 {
		t.Fail()
	}
}

func TestBitmath(t *testing.T) {
	test := []byte("aeio")

	result := test[0] & test[1] & test[2] & test[3]
	result2 := test[1] & test[2] & test[3]
	fmt.Println(string(result), string(result2))
}
