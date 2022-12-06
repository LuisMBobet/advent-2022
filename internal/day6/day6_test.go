package day6_test

import (
  "testing"
  "day6"
)

func TestDay6Part1Example(t *testing.T) {
  actual := day6.Part1("example.txt")
  expected := 7
  if actual != expected {
    t.Errorf("Expected %d but received %d",expected, actual)
  }
}

func TestDay6Part1(t *testing.T) {
  actual := day6.Part1("data.txt")
  expected := 1042
  if actual != expected {
    t.Errorf("Expected %d but received %d",expected, actual)
  }
}

func TestDay6Part2Example(t *testing.T) {
  actual := day6.Part2("example.txt", day6.ContainsUniqueChars26)
  expected := 19
  if actual != expected {
    t.Errorf("Expected %d but received %d",expected, actual)
  }
}

func TestDay6Part2(t *testing.T) {
  actual := day6.Part2("data.txt", day6.ContainsUniqueChars26)
  expected := 2980
  if actual != expected {
    t.Errorf("Expected %d but received %d",expected, actual)
  }
}

func BenchmarkTestDay6Part2WithUniqueChars26(b *testing.B){
  for n := 0; n < b.N; n++ {
    day6.Part2("data.txt", day6.ContainsUniqueChars26)
  }
}

func BenchmarkTestDay6Part2WithUniqueChars256(b *testing.B){
  for n := 0; n < b.N; n++ {
    day6.Part2("data.txt", day6.ContainsUniqueChars256)
  }
}
