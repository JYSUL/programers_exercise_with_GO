func solution(start int, end int) []int {
    var answer []int
    for i := start; i >= end; i-- {
        answer = append(answer, i)
    } 
    return answer
}