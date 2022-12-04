package day4_test

import (
  "testing"
  "day4"
)

func TestDay4Part1Example(t *testing.T) {
  actual := day4.Part1("example.txt")
  expected := 2
  if actual != expected {
    t.Errorf("Expected %d but received %d",expected, actual)
  }
}

func TestDay4Part1(t *testing.T) {
  actual := day4.Part1("data.txt")
  expected := 466
  if actual != expected {
    t.Errorf("Expected %d but received %d",expected, actual)
  }
}

func TestDay4Part2Example(t *testing.T) {
  actual := day4.Part2("example.txt")
  expected := 4
  if actual != expected {
    t.Errorf("Expected %d but received %d",expected, actual)
  }
}

func TestDay4Part2(t *testing.T) {
  actual := day4.Part2("data.txt")
  expected := 865
  if actual != expected {
    t.Errorf("Expected %d but received %d",expected, actual)
  }
}

func TestDay4Part2WithMutexConcurrency(t *testing.T) {
  actual := day4.Part2("data.txt")
  expected := 865
  if actual != expected {
    t.Errorf("Expected %d but received %d",expected, actual)
  }
}

func BenchmarkTestDay4Part2(b *testing.B){
  for n := 0; n < b.N; n++ {
    day4.Part2("data.txt")
  }
}
