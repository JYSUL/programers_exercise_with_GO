func solution(n int, m int, section []int) int {
    var end int = 0
    var answer int = 0
    for _, sect := range section {
        if sect > end {
            answer++
            end = sect + m - 1
        }
    }
    return answer
}