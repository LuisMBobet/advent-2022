package day7_test

import (
  "testing"
  "day7"
)

func TestDay7Part1Example(t *testing.T) {
  actual := day7.Part1("example.txt")
  expected := 95437
  if actual != expected {
    t.Errorf("Expected %d but received %d",expected, actual)
  }
}

func TestDay7Part1(t *testing.T) {
  actual := day7.Part1("data.txt")
  expected := 1334506
  if actual != expected {
    t.Errorf("Expected %d but received %d",expected, actual)
  }
}

func TestDay7Part2Example(t *testing.T) {
  actual := day7.Part2("example.txt")
  expected := 24933642
  if actual != expected {
    t.Errorf("Expected %d but received %d",expected, actual)
  }
}

func TestDay7Part2(t *testing.T) {
  actual := day7.Part2("data.txt")
  expected := 7421137
  if actual != expected {
    t.Errorf("Expected %d but received %d",expected, actual)
  }
}
