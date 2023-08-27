func solution(lottos []int, win_nums []int) []int {
    var min_count, max_count, zero_count int
    win_table := make(map[int]bool)
    for _, num := range win_nums {
        win_table[num] = true
    }
    
    for _, num := range lottos {
        if num == 0 {
            zero_count++
        } else if win_table[num] {
            min_count++
        }        
    }
    max_count = min_count + zero_count
    //0개맞춘 당첨은 1개맞춘 당첨과 같다.
    if min_count == 0 {
        min_count++
    }
    if max_count == 0 {
        max_count++
    }
    
    return []int{7-max_count, 7-min_count}
}