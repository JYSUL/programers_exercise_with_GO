func solution(my_string string) []int {
    var answer []int
    alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
    alphabet_table := make(map[string]int)
    for _, i := range my_string {        
        alphabet_table[string(i)] += 1        
    }
    for _, i := range alphabet {
        answer = append(answer, alphabet_table[string(i)])
    }
    return answer
}