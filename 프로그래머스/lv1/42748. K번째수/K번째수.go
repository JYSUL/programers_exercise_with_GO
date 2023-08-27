import(
    //"fmt"
    "sort"
)
func solution(array []int, commands [][]int) []int {
    var answer []int
    for _, command := range commands {
        check_slice := []int{}
        check_slice = append(check_slice, array[command[0]-1: command[1]]...)
        sort.Ints(check_slice)
        answer = append(answer, check_slice[command[2]-1])
    }
    return answer
}