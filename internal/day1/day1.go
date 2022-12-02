package day1

import (
  "os"
  "strconv"
  "strings"
  "sort"
)

func Part1(fileName string) int {
  dat,err := os.ReadFile(fileName)
  if err != nil {
    panic(err)
  }
  splitString := strings.Split(string(dat), "\n")

  maximumValue := 0
  currentElfCalories := 0

  for _, line := range splitString {
    if line != "" {
      i, err := strconv.Atoi(line)
      if err != nil {
        panic(err)
      }
      currentElfCalories += i
    } else {
      currentElfCalories = 0
    }
    if maximumValue <= currentElfCalories {
      maximumValue = currentElfCalories
    }
  }
  return maximumValue
}

func Part2(fileName string) int {
  dat,err := os.ReadFile(fileName)
  if err != nil {
    panic(err)
  }
  splitString := strings.Split(string(dat) + "\n", "\n")

  maximumValues := []int{0, 0, 0, 0}
  currentElfCalories := 0

  for _, line := range splitString {
    if line != "" {
      i, err := strconv.Atoi(line)
      if err != nil {
        panic(err)
      }
      currentElfCalories += i
    } else {
      maximumValues[0] = currentElfCalories
      sort.Ints(maximumValues)
      currentElfCalories = 0
    }

  }
  return maximumValues[1] + maximumValues[2] + maximumValues[3]
}
