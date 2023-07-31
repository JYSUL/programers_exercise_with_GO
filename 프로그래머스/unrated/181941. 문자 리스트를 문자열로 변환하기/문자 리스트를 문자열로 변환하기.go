func solution(arr []string) string {
    var s1 string
    var p1 *string = &s1
    for _, str := range(arr) {
        *p1 += str
    }
    return s1
}