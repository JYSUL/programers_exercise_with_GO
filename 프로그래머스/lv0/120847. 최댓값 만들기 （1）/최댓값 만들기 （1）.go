import (
    //"fmt"
    "sort"
)
func solution(numbers []int) int {
    sort.Ints(numbers)
    //fmt.Println(numbers[len(numbers)-1])
    return numbers[len(numbers)-1] * numbers[len(numbers)-2]
}