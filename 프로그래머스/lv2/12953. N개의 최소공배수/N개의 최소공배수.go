
func gcd(a int, b int) int {
    for a > 0 {
        a, b = b%a, a
    }
    return b
}

func solution(arr []int) int {
    var answer int = 1
    for _, num := range arr {
        answer *= num / gcd(answer, num)
    }
    
    return answer
}