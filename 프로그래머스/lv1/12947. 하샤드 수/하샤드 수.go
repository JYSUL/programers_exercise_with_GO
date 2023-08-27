func solution(x int) bool {
    var x2, sum_of_digit int = x, 0
    for x2 > 0 {
        sum_of_digit += x2 % 10
        x2 /= 10
    }
    if x % sum_of_digit == 0 {
        return true
    } else {
        return false
    }
}