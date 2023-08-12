func solution(arr []int, queries [][]int) []int {
    for _, query := range queries {
        start, end, standard_num := query[0], query[1], query[2]
        for i := start; i <= end; i++ {            
            if i % standard_num == 0 {
                arr[i] = arr[i] + 1
            }
        }
    }
    return arr
}