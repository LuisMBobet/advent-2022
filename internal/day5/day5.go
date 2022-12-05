package day5

import (
  "os"
  "strings"
  "strconv"
)

func parseStack(line string) []string{
  spaceSeparatorCount := 3
  stacks := make([]string, 0)
  for _, char := range line {
    if char == '[' || char == ']' {
    } else if char == ' ' && spaceSeparatorCount > 0 {
      spaceSeparatorCount--
    } else if char == ' ' && spaceSeparatorCount == 0  {
      stacks = append(stacks, "")
      spaceSeparatorCount = 3
    } else {
      stacks = append(stacks, string(char))
      spaceSeparatorCount = 3
    }
  }
  return stacks
}

func rotateStacks(stacks [][]string) [][]string {
  rotatedStacks := make([][]string, 0)

  for i, _ := range stacks[0] {
    tempStack := make([]string, 0)
    for _, row := range stacks {
      if row[i] != "" {
        tempStack = append([]string{row[i]},tempStack...)
      }
    }
    rotatedStacks = append(rotatedStacks, tempStack)
  }
  return rotatedStacks
}

func moveCrates(stacks [][]string, instruction string, reverseFlag bool) {
  splitInstruction := strings.Split(instruction, " ")
  numberOfCrates, err := strconv.Atoi(splitInstruction[1])
  if err != nil {
    panic(err)
  }

  position1, err := strconv.Atoi(splitInstruction[3])
  if err != nil {
    panic(err)
  }
  position1--

  position2, err := strconv.Atoi(splitInstruction[5])
  if err != nil {
    panic(err)
  }
  position2--

  position1Length := len(stacks[position1])
  cratesToMove :=  stacks[position1][position1Length-numberOfCrates:position1Length]

  if reverseFlag {
    for i, j := 0, len(cratesToMove)-1; i < j; i, j = i+1, j-1 {
      cratesToMove[i], cratesToMove[j] = cratesToMove[j], cratesToMove[i]
    }
  }

  stacks[position2] = append(stacks[position2],cratesToMove...)
  stacks[position1] = stacks[position1][0:position1Length-numberOfCrates]
}

func Part1(fileName string) string{
  dat,err := os.ReadFile(fileName)
  if err != nil {
    panic(err)
  }

  splitString := strings.Split(string(dat), "\n")

  a := make([][]string, 0)
  for i := range a {
      a[i] = make([]string, 0)
  }

  stacks := make([][]string, 0)
  var instructionsLineNumber int
  for index, line := range splitString[0:len(splitString)-1] {
    if _, err := strconv.Atoi(string(line[1])); err != nil {
      stacks = append(stacks, parseStack(line))
    } else {
      instructionsLineNumber = index + 2
      break
    }
  }
  stacks = rotateStacks(stacks)
  for _, line := range splitString[instructionsLineNumber:len(splitString)-1] {
    moveCrates(stacks,line, true)
  }

  var topRowOfCrates strings.Builder
  for _, stack := range stacks {
    topRowOfCrates.WriteString(stack[len(stack)-1])
  }
  return topRowOfCrates.String()
}

func Part2(fileName string) string{
  dat,err := os.ReadFile(fileName)
  if err != nil {
    panic(err)
  }

  splitString := strings.Split(string(dat), "\n")

  a := make([][]string, 0)
  for i := range a {
      a[i] = make([]string, 0)
  }

  stacks := make([][]string, 0)
  var instructionsLineNumber int
  for index, line := range splitString[0:len(splitString)-1] {
    if _, err := strconv.Atoi(string(line[1])); err != nil {
      stacks = append(stacks, parseStack(line))
    } else {
      instructionsLineNumber = index + 2
      break
    }
  }
  stacks = rotateStacks(stacks)
  for _, line := range splitString[instructionsLineNumber:len(splitString)-1] {
    moveCrates(stacks,line, false)
  }

  var topRowOfCrates strings.Builder
  for _, stack := range stacks {
    topRowOfCrates.WriteString(stack[len(stack)-1])
  }
  return topRowOfCrates.String()
}
