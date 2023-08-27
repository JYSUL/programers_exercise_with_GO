func solution(arr []int) float64 {
    var sum float64
    for _, num := range arr {
        sum += float64(num)
    }
    return sum / float64(len(arr))
}