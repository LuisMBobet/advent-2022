package day2

import (
  "os"
  "strings"
)

var part1PointMap = map[string]int {
  "A X": 4,
  "A Y": 8,
  "A Z": 3,
  "B X": 1,
  "B Y": 5,
  "B Z": 9,
  "C X": 7,
  "C Y": 2,
  "C Z": 6,
}


func Part1(fileName string) int{
  dat,err := os.ReadFile(fileName)
  if err != nil {
    panic(err)
  }
  splitString := strings.Split(string(dat), "\n")
  finalScore := 0
  for _, line := range splitString {
    finalScore += part1PointMap[line]
  }
  return finalScore
}

var part2PointMap = map[string]int {
  "A X": 3,
  "A Y": 4,
  "A Z": 8,
  "B X": 1,
  "B Y": 5,
  "B Z": 9,
  "C X": 2,
  "C Y": 6,
  "C Z": 7,
}

func Part2(fileName string) int {
  dat,err := os.ReadFile(fileName)
  if err != nil {
    panic(err)
  }
  splitString := strings.Split(string(dat), "\n")
  finalScore := 0
  for _, line := range splitString {
    finalScore += part2PointMap[line]
  }
  return finalScore
}
