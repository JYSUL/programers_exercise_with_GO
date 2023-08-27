
func findGCD(n int, m int) int {
    for n > 0 {
        n, m = m % n, n
    }
    return m
}

func solution(n int, m int) []int {
    return []int{findGCD(n, m), n * m / findGCD(n,m)}
}