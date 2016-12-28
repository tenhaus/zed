package main

import (
	"io/ioutil"
	"math"
	"testing"
)

func TestCompress(t *testing.T) {
	file, _ := ioutil.ReadFile("files/1000.txt")
	Compress(file)
}

func TestFill(t *testing.T) {
	var processor Processor

	processor.Points = append(processor.Points, 0x01)
	processor.Points = append(processor.Points, 0x01)
	processor.Points = append(processor.Points, 0x01)

	FillPartialProcessor(&processor)

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

func TestMapCommons(t *testing.T) {
	data := []byte{
		0x1,
		0x3,
		0x3,
		0x3,
		0x2,
		0x2,
	}

	commons := MapCommons(data)
	if commons.Keys[0] != 0x1 && commons.Keys[1] != 0x2 && commons.Keys[2] != 0x3 {
		t.Fail()
	}
}

func TestPartitionsAreOdd(t *testing.T) {

	file, _ := ioutil.ReadFile("files/octopus.jpg")
	var layer Layer
	Partition(file, &layer)

	if math.Mod(float64(len(layer.Processors)), 2.0) <= 0.0 {
		t.Fail()
	}

}

func TestGetLayerSize(t *testing.T) {
	file, _ := ioutil.ReadFile("files/20chars.txt")
	layerSize := GetLayerSize(file)

	if layerSize != 4 {
		t.Fail()
	}
}

func TestGetGridSize(t *testing.T) {
	file, _ := ioutil.ReadFile("files/5processors.txt")
	var layer Layer
	Partition(file, &layer)
	xMax, yMax := GetGridSize(len(layer.Processors))

	if xMax != 30 && yMax != 30 {
		t.Fail()
	}
}
