func solution(strArr []string) int {
    var answer int
    count_table := make(map[int]int)
    for _, str := range strArr {
        count_table[len(str)]++
    }
    for _, value := range count_table {
        if value > answer {
            answer = value
        }
    }
    return answer
    
}