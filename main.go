package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

//MaxUint maximum integer stored for minimum seeking
const MaxInt = int(^uint(0) >> 1)

//WeatherData which is Data
type WeatherData struct {
	columnName string
	compareOne int
	compareTwo int
}

func main() {
	// datas := []Data{WeatherData{}, FootballData{}}
	fmt.Println("Minimum weather data:", GetMinimumDiff())
}

//GetMinimumDiff gathers data from file to fill up Columns.
func GetMinimumDiff() WeatherData {
	wd := WeatherData{}
	minimum := MaxInt
	readLines := ReadFile("weather.dat")
	for _, value := range readLines {
		valueArrays := strings.Split(value, " ")
		valueArrays = cleanUp(valueArrays)
		name := valueArrays[0]
		if name == "Dy" || name == "mo" {
			continue
		}
		trimmedFirst, _ := strconv.Atoi(strings.TrimSuffix(valueArrays[1], "*"))
		trimmedSecond, _ := strconv.Atoi(strings.TrimSuffix(valueArrays[2], "*"))
		if (trimmedFirst - trimmedSecond) <= minimum {
			minimum = trimmedFirst - trimmedSecond
			wd.columnName = name
			wd.compareOne = trimmedFirst
			wd.compareTwo = trimmedSecond
		}
	}
	return wd
}

func (wd WeatherData) String() string {
	return fmt.Sprintf("Name: %s, Col1: %d, Col2: %d", wd.columnName, wd.compareOne, wd.compareTwo)
}

func cleanUp(s []string) (returnArr []string) {

	for _, value := range s {
		if value != "" {
			returnArr = append(returnArr, value)
		}
	}
	return
}

//ReadFile reads lines from a file and gives back a string array which contains the lines.
func ReadFile(fileName string) (fileLines []string) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if len(scanner.Text()) > 0 {
			fileLines = append(fileLines, scanner.Text())
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return
}
