func solution(arr []int, queries [][]int) []int {
    for _, query := range queries {
        start, end, standard_num := query[0], query[1], query[2]
        i := start / standard_num
        if start % standard_num != 0 {
            i++
        }
        i *= standard_num
        for ; i < end+1; i += standard_num {  
            arr[i] = arr[i] + 1            
        }
    }
    return arr
}