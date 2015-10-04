package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

//MaxUint maximum integer stored for minimum seeking
const MaxFloat = math.MaxFloat64

//WeatherData which is Data
type WeatherData struct {
	columnName string
	compareOne float64
	compareTwo float64
}

func main() {
	// datas := []Data{WeatherData{}, FootballData{}}
	fmt.Println("Minimum weather data:", GetMinimumDiff())
}

//GetMinimumDiff gathers data from file to fill up Columns.
func GetMinimumDiff() WeatherData {
	wd := WeatherData{}
	minimum := MaxFloat
	readLines := ReadFile("weather.dat")
	for _, value := range readLines {
		valueArrays := strings.Split(value, ",")
		name := valueArrays[0]
		trimmedFirst, err := strconv.ParseFloat(strings.TrimSuffix(valueArrays[1], "*"), 64)
		if err != nil {
			continue
		}
		trimmedSecond, _ := strconv.ParseFloat(strings.TrimSuffix(valueArrays[2], "*"), 64)
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
	return fmt.Sprintf("Name: %s, Col1: %f, Col2: %f", wd.columnName, wd.compareOne, wd.compareTwo)
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
		line := scanner.Text()
		if len(line) > 0 {
			re := regexp.MustCompile("\\w+")
			lines := re.FindAllString(line, -1)
			fileLines = append(fileLines, strings.Join(lines, ","))
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return
}
