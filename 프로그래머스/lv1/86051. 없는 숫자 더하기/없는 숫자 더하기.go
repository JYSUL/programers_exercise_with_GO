func solution(numbers []int) int {
    var answer int = 0
    check_table := make(map[int]bool)
    for _, number := range numbers {
        check_table[number] = true
    }
    for i := 0; i < 10; i++ {
        if !check_table[i] {
            answer += i
        }
    }
    return answer
}