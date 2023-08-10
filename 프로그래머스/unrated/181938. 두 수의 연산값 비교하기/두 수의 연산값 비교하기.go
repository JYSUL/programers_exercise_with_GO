import (
    "math"
    "strconv"
)
func solution(a int, b int) int {
    cal_num, _ := strconv.Atoi(strconv.Itoa(a) + strconv.Itoa(b))
    return int(math.Max(float64(cal_num),float64(2 * a * b)))
}