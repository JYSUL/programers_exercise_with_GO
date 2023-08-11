func solution(a int, b int, c int) int {
    if a == b && b == c {
        return (a*a*a + b*b*b + c*c*c) * (a * a + b * b + c * c ) * (a + b + c)
    } else if a == b || b == c || c == a {
        return (a * a + b * b + c * c ) * (a + b + c)
    } else {
        return a + b + c
    }
    return 0
}