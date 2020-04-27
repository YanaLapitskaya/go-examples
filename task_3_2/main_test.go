package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var equalMsg = "Values should be equal"

func TestSquareArea(t *testing.T) {
	testCases := []struct {
		in  float64
		out float64
	}{
		{5, 25},
		{1, 1},
		{500, 250000},
	}
	for _, value := range testCases {
		var s Figure = Square{value.in}
		var result = s.area()
		assert.Equal(t, result, value.out, equalMsg)
	}
}

func TestSquarePerimeter(t *testing.T) {
	testCases := []struct {
		in  float64
		out float64
	}{
		{5, 20},
		{1, 4},
		{500, 2000},
	}
	for _, value := range testCases {
		var s Figure = Square{value.in}
		var result = s.perimeter()
		assert.Equal(t, result, value.out, equalMsg)
	}
}

func TestCircleArea(t *testing.T) {
	testCases := []struct {
		in  float64
		out float64
	}{
		{5, 78.53981633974483},
		{1, 3.141592653589793},
		{500, 785398.1633974483},
	}
	for _, value := range testCases {
		var s Figure = Circle{value.in}
		var result = s.area()
		assert.Equal(t, result, value.out, equalMsg)
	}
}

func TestCirclePerimeter(t *testing.T) {
	testCases := []struct {
		in  float64
		out float64
	}{
		{5, 31.41592653589793},
		{1, 6.283185307179586},
		{500, 3141.592653589793},
	}
	for _, value := range testCases {
		var s Figure = Circle{value.in}
		var result = s.perimeter()
		assert.Equal(t, result, value.out, equalMsg)
	}
}
