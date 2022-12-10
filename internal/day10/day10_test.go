package day10_test

import (
  "testing"
  "day10"
)

func TestDay10Part1Example(t *testing.T) {
  actual := day10.Part1("example.txt")
  expected := 13140
  if actual != expected {
    t.Errorf("Expected %d but received %d",expected, actual)
  }
}

func TestDay10Part1(t *testing.T) {
  actual := day10.Part1("data.txt")
  expected := 1807
  if actual != expected {
    t.Errorf("Expected %d but received %d",expected, actual)
  }
}

// func TestDay10Part2Example(t *testing.T) {
//   actual := day10.Part2("example.txt")
//   expected := 8
//   if actual != expected {
//     t.Errorf("Expected %d but received %d",expected, actual)
//   }
// }
//
// func TestDay10Part2(t *testing.T) {
//   actual := day10.Part2("data.txt")
//   expected := 480000
//   if actual != expected {
//     t.Errorf("Expected %d but received %d",expected, actual)
//   }
// }
