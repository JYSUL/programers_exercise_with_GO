import "math"
func solution(n int64) int64 {
    if int64(math.Sqrt(float64(n))) * int64(math.Sqrt(float64(n))) == n {
        return (int64(math.Sqrt(float64(n))) + 1) * (int64(math.Sqrt(float64(n))) + 1)
    } else {
        return -1
    }
}