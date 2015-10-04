package main

import (
	"fmt"
	"testing"
)

//TestAbleToReadFromFile tests the ability to read from a file line by line.
func TestAbleToReadFromFile(t *testing.T) {
	var fileLines []string
	if fileLines = ReadFile("weather.dat"); len(fileLines) == 0 {
		t.Error("Failed to read lines from football.dat.")
	}
	for _, value := range fileLines {
		fmt.Println(value)
	}
}
