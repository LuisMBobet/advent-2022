package day3_test

import (
  "testing"
  "day3"
)

func TestDay3Part1Example(t *testing.T) {
  actual := day3.Part1("example.txt")
  expected := 157
  if actual != expected {
    t.Errorf("Expected %d but received %d",expected, actual)
  }
}

func TestDay3Part1(t *testing.T) {
  actual := day3.Part1("data.txt")
  expected := 7878
  if actual != expected {
    t.Errorf("Expected %d but received %d",expected, actual)
  }
}

func TestDay3Part2Example(t *testing.T) {
  actual := day3.Part2("example.txt")
  expected := 70
  if actual != expected {
    t.Errorf("Expected %d but received %d",expected, actual)
  }
}

func TestDay3Part2(t *testing.T) {
  actual := day3.Part2("data.txt")
  expected := 2760
  if actual != expected {
    t.Errorf("Expected %d but received %d",expected, actual)
  }
}
