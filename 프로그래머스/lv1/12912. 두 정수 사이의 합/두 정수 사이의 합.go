func solution(a int, b int) int64 {
    var answer int64
    if a > b {
        a, b = b, a
    }
    for i:=a; i<= b; i++ {
        answer += int64(i)
    }
    return answer
}