import (
    //"fmt"
    //"sort"
)
func solution(numbers []int) int {
    /*
    sort.Ints(numbers)
    //fmt.Println(numbers[len(numbers)-1])
    return numbers[len(numbers)-1] * numbers[len(numbers)-2]
    */
    var max int = 0
    var second_max int = 0
    for i := 0; i < len(numbers); i++ {
        if numbers[i] > max {
            second_max = max
            max = numbers[i]
        } else if numbers[i] > second_max{
            second_max = numbers[i]
        }
    }
    return max * second_max
}