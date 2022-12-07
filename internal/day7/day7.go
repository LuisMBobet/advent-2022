package day7

import (
  "os"
  "strings"
  "fmt"
  "strconv"
)

func contains[T comparable](s []T, e T) bool {
    for _, v := range s {
        if v == e {
            return true
        }
    }
    return false
}

func calcSumOfDirSizesLessThanCriteria (dirSizeMap map[string]int, criteria int) int {
  var sum int
  for _, element := range dirSizeMap {
    if criteria >= element {
      sum += element
    }
  }
  return sum
}

func Part1(fileName string) int{
  dat,err := os.ReadFile(fileName)
  if err != nil {
    panic(err)
  }

  splitString := strings.Split(string(dat), "\n")

  dirs := make(map[string]int)
  var currentDir []string
  var searchedPaths []string
  var newDir bool
  for _, line := range splitString[0: len(splitString)-1] {
    splitLine := strings.Split(line, " ")
    currentDirString := strings.Join(currentDir, "/")
    if splitLine[1] == "cd" && splitLine[2] != ".." {
      currentDir = append(currentDir, splitLine[2])
      newDir = false
    } else if splitLine[1] == "cd" && splitLine[2] == ".." {
      currentDir = currentDir[0:len(currentDir)-1]
      newDir = false
    } else if splitLine[1] == "ls" && !contains(searchedPaths, currentDirString){
      searchedPaths = append(searchedPaths, currentDirString)
      newDir = true
    } else if newDir {
      intFirstValue, err := strconv.Atoi(splitLine[0])
      if err == nil {
        for index, _ := range currentDir {
          fullFilePath := strings.Join(currentDir[0:index+1], "/")
          dirs[fullFilePath] = dirs[fullFilePath] + intFirstValue
        }
      }
    }
  }
  fmt.Println(dirs)
  return calcSumOfDirSizesLessThanCriteria(dirs, 100000)
}
