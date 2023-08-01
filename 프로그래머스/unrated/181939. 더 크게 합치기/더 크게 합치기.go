import "strconv"
func solution(a int, b int) int {
    var a_first int
    var b_first int
    a_first, _ = strconv.Atoi(strconv.Itoa(a) + strconv.Itoa(b))
    b_first, _ = strconv.Atoi(strconv.Itoa(b) + strconv.Itoa(a))
    if a_first > b_first {
        return a_first
    } else {
        return b_first
    }
}