import (
    "sort"
)
func solution(k int, tangerines []int) int {
    var result int = 0
    tangerine_table := make(map[int]int)
    var count_slice []int
    
    for _, tangerine := range tangerines {
        tangerine_table[tangerine]++
    }
    
    for _, value := range tangerine_table {
        count_slice = append(count_slice, value)
    }
    sort.Sort(sort.Reverse(sort.IntSlice(count_slice)))
    
    for _, value := range count_slice {
        k -= value
        result++
        if k <= 0 {
            break
        }
    }
    return result
}