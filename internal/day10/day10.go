package day10

import (
  "os"
  "strings"
  "strconv"
)

func readFileIntoLines(filePath string) []string {
  dat,err := os.ReadFile(filePath)
  if err != nil {
    panic(err)
  }
  lines := strings.Split(string(dat), "\n")
  return lines[0:len(lines)-1]
}

func executeCommand(instruction string, cycles int, memoryValue int) (int, int){
  splitInstruction := strings.Split(instruction, " ")
  if splitInstruction[0] == "addx" {
    cycles += 2
    numberToAdd, err := strconv.Atoi(splitInstruction[1])
    if err != nil {
      panic(err)
    }
    memoryValue += numberToAdd
  } else {
    cycles += 1
  }
  return cycles, memoryValue

}

func Part1(fileName string) int{
  instructions := readFileIntoLines(fileName)
  registerXValue, cycles := 1, 0
  interestingSignals := []int{20,60,100,140,180,220}
  interestingSignalStrengthsSum := 0
  for _, instruction := range instructions {
    tempRegisterValue := registerXValue
    cycles,registerXValue = executeCommand(instruction, cycles, registerXValue)
    if cycles >= interestingSignals[0] {
      interestingSignalStrengthsSum += interestingSignals[0] * tempRegisterValue

      interestingSignals = interestingSignals[1:len(interestingSignals)]
      if len(interestingSignals) == 0 {
        break
      }
    }
  }
  return interestingSignalStrengthsSum
}
