package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestCompress(t *testing.T) {
	file, _ := ioutil.ReadFile("files/optimal.txt")

	if out, err := os.Create("out.zed"); err != nil {
		t.Fail()
	} else {
		Compress(file, out)
	}
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

func TestPartition(t *testing.T) {
	file, _ := ioutil.ReadFile("files/double.txt")
	var layer Layer
	Partition(file, &layer)

	if len(layer.Processors) != 2 {
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

func TestSliceCommons(t *testing.T) {
	file, _ := ioutil.ReadFile("files/double.txt")
	commons := MapCommons(file)

	if len(commons.Slice()) != 2 {
		t.Fail()
	}

	uneven, _ := ioutil.ReadFile("files/double_uneven.txt")
	commons2 := MapCommons(uneven)

	if len(commons2.Slice()) != 2 {
		t.Fail()
	}
}
