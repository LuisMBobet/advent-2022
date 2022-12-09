package day8_test

import (
  "testing"
  "day8"
)

func TestDay8Part1Example(t *testing.T) {
  actual := day8.Part1("example.txt")
  expected := 21
  if actual != expected {
    t.Errorf("Expected %d but received %d",expected, actual)
  }
}

func TestDay8Part1(t *testing.T) {
  actual := day8.Part1("data.txt")
  expected := 1807
  if actual != expected {
    t.Errorf("Expected %d but received %d",expected, actual)
  }
}

func TestDay8Part2Example(t *testing.T) {
  actual := day8.Part2("example.txt")
  expected := 8
  if actual != expected {
    t.Errorf("Expected %d but received %d",expected, actual)
  }
}

func TestDay8Part2(t *testing.T) {
  actual := day8.Part2("data.txt")
  expected := 480000
  if actual != expected {
    t.Errorf("Expected %d but received %d",expected, actual)
  }
}
