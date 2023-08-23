func solution(a int, b int, n int) int {
    var answer int = 0
    for n >= a {
        answer += (n / a) * b
        n = (n / a) * b + n % a
    }
    return answer
}

/*
answer = 0
    while n >= a :
        answer += (n//a) * b
        n = (n // a) * b + n % a    

*/