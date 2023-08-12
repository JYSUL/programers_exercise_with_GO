func solution(l int, r int) []int {
    var answer []int
    var flag bool
    for i := l; i <= r; i++ {
        check_num := i
        flag = true
        for check_num > 0 {
            if check_num % 10 == 5 || check_num % 10 == 0 {
                check_num /= 10
            } else {
                flag = false
                break
            }
        }
        if flag {
            answer = append(answer, i)
        }
    }
    if len(answer) == 0 {
        answer = append(answer, -1)
    }
    return answer
}