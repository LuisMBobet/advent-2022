package day6

import (
  "os"
  "strings"
)

type fn func(string) bool

func ContainsUniqueChars256(input string) bool {
  var alphabet [256]bool
  for _, ascii := range input {

    if alphabet[ascii] {
      return false
    }
    alphabet[ascii] = true
  }

  return true
}

func ContainsUniqueChars26(input string) bool {
  var alphabet [26]bool
  for _, ascii := range input {
    ascii = ascii - 97
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
    if ContainsUniqueChars26(splitString[0][i:i+4]) {
      return i+4
    }
  }
  panic("No unique characters found")
}

func Part2(fileName string, uniqueCharFunc fn) int{
  dat,err := os.ReadFile(fileName)
  if err != nil {
    panic(err)
  }

  splitString := strings.Split(string(dat), "\n")

  for i, _ := range splitString[0] {
    if uniqueCharFunc(splitString[0][i:i+14]) {
      return i+14
    }
  }
  panic("No unique characters found")
}
