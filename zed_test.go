package main

import (
	"io/ioutil"
	"testing"
)

func TestCompress(t *testing.T) {
	file, _ := ioutil.ReadFile("files/simple.txt")
	Compress(file)
}
