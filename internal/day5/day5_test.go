package day5_test

import (
  "testing"
  "day5"
)

func TestDay5Part1Example(t *testing.T) {
  actual := day5.Part1("example.txt")
  expected := "CMZ"
  if actual != expected {
    t.Errorf("Expected %s but received %s",expected, actual)
  }
}

func TestDay5Part1(t *testing.T) {
  actual := day5.Part1("data.txt")
  expected := "SVFDLGLWV"
  if actual != expected {
    t.Errorf("Expected %s but received %s",expected, actual)
  }
}

func TestDay5Part2Example(t *testing.T) {
  actual := day5.Part2("example.txt")
  expected := "MCD"
  if actual != expected {
    t.Errorf("Expected %s but received %s",expected, actual)
  }
}

func TestDay5Part2(t *testing.T) {
  actual := day5.Part2("data.txt")
  expected := "DCVTCVPCL"
  if actual != expected {
    t.Errorf("Expected %s but received %s",expected, actual)
  }
}
