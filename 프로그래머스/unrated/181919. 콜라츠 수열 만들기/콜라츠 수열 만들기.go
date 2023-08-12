func solution(n int) []int {
    var answer []int
    for n != 1 {
        answer = append(answer, n)
        if n % 2 == 0 {
            n = n / 2
        } else {
            n = n * 3 + 1
        }
    }
    answer = append(answer, 1)
    return answer
}