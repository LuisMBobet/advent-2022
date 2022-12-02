package day1_test

import (
  "testing"
  "day1"
)

func TestDay1Part1Example(t *testing.T) {
  actual := day1.Part1("example.txt")
  expected := 24000
  if actual != expected {
    t.Errorf("Expected %d but received %d",expected, actual)
  }
}

func TestDay1Part1(t *testing.T) {
  actual := day1.Part1("data.txt")
  expected := 67450
  if actual != expected {
    t.Errorf("Expected %d but received %d",expected, actual)
  }
}

func TestDay1Part2Example(t *testing.T) {
  actual := day1.Part2("example.txt")
  expected := 45000
  if actual != expected {
    t.Errorf("Expected %d but received %d",expected, actual)
  }
}

func TestDay1Part2(t *testing.T) {
  actual := day1.Part2("data.txt")
  expected := 199357
  if actual != expected {
    t.Errorf("Expected %d but received %d",expected, actual)
  }
}
