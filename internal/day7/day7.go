package day7

import (
  "os"
  "strings"
  "strconv"
)

func calcSumOfDirSizesLessThanCriteria (dirSizeMap map[string]int, criteria int) int {
  var sum int
  for _, element := range dirSizeMap {
    if criteria >= element {
      sum += element
    }
  }
  return sum
}

func findSmallestDirLargerThanCriteria (dirSizeMap map[string]int, criteria int) int {
  smallestDirSize := int(^uint(0)  >> 1)
  for _, element := range dirSizeMap {
    if element >= criteria && element <= smallestDirSize{
      smallestDirSize = element
    }
  }
  return smallestDirSize
}

func Part1(fileName string) int{
  dat,err := os.ReadFile(fileName)
  if err != nil {
    panic(err)
  }

  splitString := strings.Split(string(dat), "\n")

  dirs := make(map[string]int)
  var currentDir []string
  for _, line := range splitString[0: len(splitString)-1] {
    splitLine := strings.Split(line, " ")
    if splitLine[1] == "cd" && splitLine[2] != ".." {
      currentDir = append(currentDir, splitLine[2])
    } else if splitLine[1] == "cd" && splitLine[2] == ".." {
      currentDir = currentDir[0:len(currentDir)-1]
    } else {
      intFirstValue, err := strconv.Atoi(splitLine[0])
      if err == nil {
        for index, _ := range currentDir {
          fullFilePath := strings.Join(currentDir[0:index+1], "/")
          dirs[fullFilePath] += intFirstValue
        }
      }
    }
  }
  return calcSumOfDirSizesLessThanCriteria(dirs, 100000)
}


func Part2(fileName string) int{
  dat,err := os.ReadFile(fileName)
  if err != nil {
    panic(err)
  }

  splitString := strings.Split(string(dat), "\n")

  dirs := make(map[string]int)
  var currentDir []string
  for _, line := range splitString[0: len(splitString)-1] {
    splitLine := strings.Split(line, " ")
    if splitLine[1] == "cd" && splitLine[2] != ".." {
      currentDir = append(currentDir, splitLine[2])
    } else if splitLine[1] == "cd" && splitLine[2] == ".." {
      currentDir = currentDir[0:len(currentDir)-1]
    } else {
      intFirstValue, err := strconv.Atoi(splitLine[0])
      if err == nil {
        for index, _ := range currentDir {
          fullFilePath := strings.Join(currentDir[0:index+1], "/")
          dirs[fullFilePath] += intFirstValue
        }
      }
    }
  }
  requiredSpace := 30000000 - (70000000 - dirs["/"])
  return findSmallestDirLargerThanCriteria(dirs, requiredSpace)
}
