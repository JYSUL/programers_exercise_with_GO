import "strconv"
func solution(s string) int {
    num, err := strconv.Atoi(s)
    if err != nil {
        panic("panic")
    }
    return num
}