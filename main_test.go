package main

import (
	"bufio"
	"strings"
	"testing"
)

var testData = "1 2 3 3 4 4 5 6 7 7"
var expectedData = []int {1 ,2 ,3, 4, 5, 6 ,7}

func TestNativeParse(t *testing.T)  {

	in := bufio.NewReader(strings.NewReader(testData))

	result, err := nativeParseAndRemoveDuplicates(in)

	if err != nil {
		t.Errorf("native parse failed with error")
	}

	for pos, number := range result {
		if number != expectedData[pos] {
			t.Errorf("native parse failed with error - not matching %v %v", expectedData, result )
		}
	}

}

func TestCustomParse(t *testing.T) {
	in := bufio.NewReader(strings.NewReader(testData))

	result, err := customParseAndRemoveDuplicates(in)

	if err != nil {
		t.Errorf("custom parse failed with error")
	}

	for pos, number := range result {
		if number != expectedData[pos] {
			t.Errorf("custom parse failed with error - not matching %v %v", expectedData, result )
		}
	}
}