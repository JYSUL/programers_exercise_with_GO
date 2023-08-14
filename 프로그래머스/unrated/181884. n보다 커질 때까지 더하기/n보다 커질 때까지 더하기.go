func solution(numbers []int, n int) int {
    var answer, index int = 0, 0
    for answer <= n {
        answer += numbers[index]
        index++
    }
    return answer
}