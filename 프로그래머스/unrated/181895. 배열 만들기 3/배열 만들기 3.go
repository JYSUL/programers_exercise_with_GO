func solution(arr []int, intervals [][]int) []int {
    var answer []int
    answer = append(answer, arr[intervals[0][0]:intervals[0][1]+1]...)
    answer = append(answer, arr[intervals[1][0]:intervals[1][1]+1]...)
    return answer
}