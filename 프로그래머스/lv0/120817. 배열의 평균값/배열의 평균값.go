func solution(numbers []int) float64 {
    var total_sum float64 = 0
    for _, score := range numbers {
        total_sum = total_sum + float64(score)
    }
    return total_sum / float64(len(numbers))
}