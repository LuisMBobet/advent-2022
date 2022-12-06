package day6

import (
  "os"
  "strings"
)

func containsUniqueChars(input string) bool {
  var alphabet [256]bool
  for _, ascii := range input {

    if alphabet[ascii] {
      return false
    }
    alphabet[ascii] = true
  }

  return true
}

func Part1(fileName string) int{
  dat,err := os.ReadFile(fileName)
  if err != nil {
    panic(err)
  }

  splitString := strings.Split(string(dat), "\n")

  for i, _ := range splitString[0] {
    if containsUniqueChars(splitString[0][i:i+4]) {
      return i+4
    }
  }
  panic("No unique characters found")
}

func Part2(fileName string) int{
  dat,err := os.ReadFile(fileName)
  if err != nil {
    panic(err)
  }

  splitString := strings.Split(string(dat), "\n")

  for i, _ := range splitString[0] {
    if containsUniqueChars(splitString[0][i:i+14]) {
      return i+14
    }
  }
  panic("No unique characters found")
}
