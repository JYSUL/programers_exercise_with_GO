import "math"
func solution(left int, right int) int {
    var answer int = 0
    for num := left; num <= right; num++ {
        if int(math.Sqrt(float64(num))) * int(math.Sqrt(float64(num))) == num {
            answer -= num
        } else {
            answer += num
        }
    }
    return answer
}