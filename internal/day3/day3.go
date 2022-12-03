package day3

import (
  "os"
  "strings"
  "unicode"
)

func contains(slice []rune, char rune) bool {
  for _, value := range slice {
    if value == char {
      return true
    }
  }
  return false
}

func calculatePriorityForItem(item rune) int {
  var asciiConstant int
  if unicode.IsUpper(item) {
    asciiConstant = 38
  } else {
    asciiConstant = 96
  }
  return int(item) - asciiConstant
}

func Part1(fileName string) int{
  dat,err := os.ReadFile(fileName)
  if err != nil {
    panic(err)
  }
  splitString := strings.Split(string(dat), "\n")
  sharedItemPriorityTotal := 0

  for _, rucksackContents := range splitString {
    length := len(rucksackContents)
    compartment1, compartment2 := rucksackContents[0:length/2], rucksackContents[length/2:length]
    var previouslySearchedItems []rune

    for _, item := range compartment1 {
      if contains(previouslySearchedItems, item){
        continue
      }

      if strings.Contains(compartment2, string(item)) {
        sharedItemPriorityTotal += calculatePriorityForItem(item)
        previouslySearchedItems = append(previouslySearchedItems, item)
        continue
      }
    }
  }
  return sharedItemPriorityTotal
}

func Part2(fileName string) int {
  dat,err := os.ReadFile(fileName)
  if err != nil {
    panic(err)
  }

  splitString := strings.Split(string(dat), "\n")
  sharedItemPriorityTotal := 0


  for index := 0; index < len(splitString) -1; index+=3 {
    firstRucksackContents := splitString[index]
    secondRucksackContents := splitString[index+1]
    thirdRucksackContents := splitString[index+2]

    var previouslySearchedItems []rune

    for _, item := range firstRucksackContents {
      if contains(previouslySearchedItems, item){
        continue
      }

      if strings.Contains(secondRucksackContents, string(item)){
        if strings.Contains(thirdRucksackContents, string(item)){
          sharedItemPriorityTotal += calculatePriorityForItem(item)
          previouslySearchedItems = append(previouslySearchedItems, item)
          continue
        }
      }
    }
  }
  return sharedItemPriorityTotal
}
