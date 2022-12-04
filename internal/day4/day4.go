package day4

import (
  "os"
  "strings"
  "strconv"
  "sync"
)

func Part1(fileName string) int{
  dat,err := os.ReadFile(fileName)
  if err != nil {
    panic(err)
  }

  splitString := strings.Split(string(dat), "\n")
  fullyContainedSections := 0

  for _, line := range splitString[0:len(splitString)-1] {
    var intPairs []int
    pairs := strings.Split(strings.Replace(line, ",","-", 1), "-")

    for _, pair := range pairs {
      intPair, err := strconv.Atoi(pair)
      if err != nil {
        panic(err)
      }
      intPairs = append(intPairs, intPair)
    }

    if (intPairs[0] >= intPairs[2] && intPairs[1] <= intPairs[3]) || (intPairs[0] <= intPairs[2] && intPairs[1] >= intPairs[3]){
      fullyContainedSections++
    }
  }

  return fullyContainedSections
}

func Part2(fileName string) int{
  dat,err := os.ReadFile(fileName)
  if err != nil {
    panic(err)
  }

  splitString := strings.Split(string(dat), "\n")
  fullyContainedSections := 0

  for _, line := range splitString[0:len(splitString)-1] {
    var intPairs []int
    pairs := strings.Split(strings.Replace(line, ",","-", 1), "-")

    for _, pair := range pairs {
      intPair, err := strconv.Atoi(pair)
      if err != nil {
        panic(err)
      }
      intPairs = append(intPairs, intPair)
    }

    if (intPairs[0] >= intPairs[2] && intPairs[0] <= intPairs[3]) ||
    (intPairs[1] >= intPairs[2] && intPairs[1] <= intPairs[3]) ||
    (intPairs[2] >= intPairs[0] && intPairs[2] <= intPairs[1]) ||
    (intPairs[3] >= intPairs[0] && intPairs[3] <= intPairs[1]){
      fullyContainedSections++
    }
  }

  return fullyContainedSections
}

func Part2WithMutexConcurrency(fileName string) int{
  dat,err := os.ReadFile(fileName)
  if err != nil {
    panic(err)
  }

  splitString := strings.Split(string(dat), "\n")
  fullyContainedSections := 0


  var wg sync.WaitGroup
	var mtx sync.RWMutex

  for _, line := range splitString[0:len(splitString)-1] {
    wg.Add(1)
    line := line

    go func() {
      var intPairs []int
      pairs := strings.Split(strings.Replace(line, ",","-", 1), "-")

      for _, pair := range pairs {
        intPair, err := strconv.Atoi(pair)
        if err != nil {
          panic(err)
        }
        intPairs = append(intPairs, intPair)
      }

      if (intPairs[0] >= intPairs[2] && intPairs[0] <= intPairs[3]) ||
      (intPairs[1] >= intPairs[2] && intPairs[1] <= intPairs[3]) ||
      (intPairs[2] >= intPairs[0] && intPairs[2] <= intPairs[1]) ||
      (intPairs[3] >= intPairs[0] && intPairs[3] <= intPairs[1]){
        mtx.Lock()
        fullyContainedSections++
        mtx.Unlock()
      }
      wg.Done()
    }()
  }

  wg.Wait()

  return fullyContainedSections
}
