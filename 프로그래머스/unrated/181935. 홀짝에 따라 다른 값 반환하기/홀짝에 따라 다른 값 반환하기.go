import "math"
func solution(n int) int {
    var answer int = 0
    var seq int = n % 2
    if seq == 1 {
        for ; seq <=n; seq += 2 {
            answer += seq
        }
    } else {
        for ; seq <= n; seq += 2 {
            answer += int(math.Pow(float64(seq), 2))
        }
    }
    return answer
}