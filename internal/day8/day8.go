package day8

import (
  "os"
  "strings"
)

func readFileIntoLines(filePath string) []string {
  dat,err := os.ReadFile(filePath)
  if err != nil {
    panic(err)
  }
  lines := strings.Split(string(dat), "\n")
  return lines[0:len(lines)-1]
}

func rotateTrees(trees []string) []string {
  lengthOfTrees := len(trees)
  var verticalTrees []string
  for index, _ := range trees[0:lengthOfTrees] {
    var tempTrees string
    for _, row := range trees {
      tempTrees = string(row[index]) + tempTrees
    }
    verticalTrees = append(verticalTrees, tempTrees)
  }
  return verticalTrees
}

func rotateVisibleTrees(visibleTrees [][]bool) [][]bool {
  rotatedCountedTrees := make([][]bool, 0)

  for i, _ := range visibleTrees[0] {
    tempSlice := make([]bool, 0)
    for _, row := range visibleTrees {
      tempSlice = append([]bool{row[i]},tempSlice...)
    }
    rotatedCountedTrees = append(rotatedCountedTrees, tempSlice)
  }
  return rotatedCountedTrees
}

func rotateScenicTrees(visibleTrees [][]int) [][]int {
  rotatedCountedTrees := make([][]int, 0)

  for i, _ := range visibleTrees[0] {
    tempSlice := make([]int, 0)
    for _, row := range visibleTrees {
      tempSlice = append([]int{row[i]},tempSlice...)
    }
    rotatedCountedTrees = append(rotatedCountedTrees, tempSlice)
  }
  return rotatedCountedTrees
}

func markVisibleTrees(trees []string, visibleTrees [][]bool) {
  for index, row := range trees {
    markedTrees := 0
    smallestTree := '/'
    for indexj, tree := range row {
      if tree > smallestTree{
        markedTrees ++
        smallestTree = tree
        visibleTrees[index][indexj] = true
      }
    }
  }
}

func calculateScenicTrees(trees []string, scenicTrees [][]int) {
  for index, row := range trees {
    for indexj, tree := range row {
      numberOfVisibleTrees := 0
      for indexk := indexj+1;indexk<=len(row)-1;indexk++ {
        numberOfVisibleTrees ++
        if tree <= rune(row[indexk]) {
          break
        }
      }
      scenicTrees[index][indexj] *= numberOfVisibleTrees
    }
  }
}

func countTrees(visibleTrees [][]bool) int {
  visibleTreeCount := 0
  for _, row := range visibleTrees {
    for _, tree := range row {
      if tree {
        visibleTreeCount += 1
      }
    }
  }
  return visibleTreeCount
}

func reverseStringSlice(slice []string) {
  for i, row := range slice {
    splitRow := []rune(row)
    for i, j := 0, len(splitRow)-1; i < j; i, j = i+1, j-1 {
        splitRow[i], splitRow[j] = splitRow[j], splitRow[i]
    }
    slice[i] = string(splitRow)
  }

}
func reverseBoolSlice(slice []bool) {
  for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
    slice[i], slice[j] = slice[j], slice[i]
  }
}

func reverseIntSlice(slice []int) {
  for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
    slice[i], slice[j] = slice[j], slice[i]
  }
}

func countAscTreesFromEnd(trees []string) int{

  visibleTrees := make([][]bool, len(trees))
  for i := range visibleTrees {
      visibleTrees[i] = make([]bool, len(trees[0]) )

  }
  markVisibleTrees(trees, visibleTrees)

  reverseStringSlice(trees)
  for _, row := range visibleTrees {
    reverseBoolSlice(row)
  }

  markVisibleTrees(trees, visibleTrees)

  trees = rotateTrees(trees)
  visibleTrees = rotateVisibleTrees(visibleTrees)
  markVisibleTrees(trees, visibleTrees)

  reverseStringSlice(trees)
  for _, row := range visibleTrees {
    reverseBoolSlice(row)
  }
  markVisibleTrees(trees, visibleTrees)

  return countTrees(visibleTrees)
}

func findMaxScore(scenicTrees [][]int) int {
  maxScore := 0
  for _, row := range scenicTrees {
    for _, col := range row {
      if col > maxScore {
        maxScore = col
      }
    }
  }
  return maxScore
}

func findMostScenicTree(trees []string) int{

  scenicTrees := make([][]int, len(trees))
  for i := range scenicTrees {
    scenicTrees[i] = make([]int, len(trees[0]))
    for j := range scenicTrees[i] {
      scenicTrees[i][j] = 1
    }
  }

  calculateScenicTrees(trees, scenicTrees)

  reverseStringSlice(trees)
  for _, row := range scenicTrees {
    reverseIntSlice(row)
  }

  calculateScenicTrees(trees, scenicTrees)

  trees = rotateTrees(trees)
  scenicTrees = rotateScenicTrees(scenicTrees)
  calculateScenicTrees(trees, scenicTrees)

  reverseStringSlice(trees)
  for _, row := range scenicTrees {
    reverseIntSlice(row)
  }
  calculateScenicTrees(trees, scenicTrees)


  return findMaxScore(scenicTrees)
}

func Part1(fileName string) int{
  trees := readFileIntoLines(fileName)

  return countAscTreesFromEnd(trees)
}

func Part2(fileName string) int{
  trees := readFileIntoLines(fileName)
  return findMostScenicTree(trees)
}
