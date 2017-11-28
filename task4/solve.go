package main
import (
  "strings"
  "unicode"
)
func RemoveEven(sx []int) []int {
  var res []int
  for _, val := range sx {
    if val % 2 != 0 {
      res = append(res, val)
    }
  }
  return res
}
func PowerGenerator(x int) (func() int) {
  val := 1
  return func() int {
    val *= x
    return val
  }
}
func DifferentWordsCount(s string) int {
  res := make(map[string]int)
  var word string
  for _, c := range s {
    if unicode.IsLetter(c) {
      word += strings.ToLower(string(c))
    } else if len(word) != 0 {
      res[word] = 1
      word = ""
    }
  }
  if len(word) != 0 {
    res[word] = 1
  }
  return len(res)
}
