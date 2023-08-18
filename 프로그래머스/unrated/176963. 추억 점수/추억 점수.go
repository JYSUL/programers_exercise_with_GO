func solution(name []string, yearning []int, photo [][]string) []int {
    var answer_slice []int
    yearning_table := make(map[string]int)
    var points int
    
    for i := 0; i < len(name); i++ {
        yearning_table[name[i]] = yearning[i]
    }
    
    for _, people := range photo {
        points = 0
        for _, person := range people {
            points += yearning_table[person]
        }
        answer_slice = append(answer_slice, points)
    }
    return answer_slice
}