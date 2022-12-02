package day2_test

import (
  "testing"
  "day2"
)

func TestDay2Part1Example(t *testing.T) {
  actual := day2.Part1("example.txt")
  expected := 15
  if actual != expected {
    t.Errorf("Expected %d but received %d",expected, actual)
  }
}

func TestDay2Part1(t *testing.T) {
  actual := day2.Part1("data.txt")
  expected := 11841
  if actual != expected {
    t.Errorf("Expected %d but received %d",expected, actual)
  }
}

func TestDay2Part2Example(t *testing.T) {
  actual := day2.Part2("example.txt")
  expected := 12
  if actual != expected {
    t.Errorf("Expected %d but received %d",expected, actual)
  }
}

func TestDay2Part2(t *testing.T) {
  actual := day2.Part2("data.txt")
  expected := 13022
  if actual != expected {
    t.Errorf("Expected %d but received %d",expected, actual)
  }
}
