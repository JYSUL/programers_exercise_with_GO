func solution(arr []int) []int {
    var answer []int
    for _, num := range arr {
        for i := 0; i < num; i++ {
            answer = append(answer, num)
        }
    }
    return answer
}