func solution(arr []int, queries [][]int) []int {
    var min_num int = -1
    var answer_slice []int
    
    for _, query := range queries {
        
        min_num = 10000000
        start, end, standard_num := query[0], query[1], query[2]
        for _, check_num := range arr[start:end+1] {
            
            if check_num > standard_num && check_num < min_num {
                min_num = check_num
            }
        }
        
        if min_num == 10000000 {
            min_num = -1
        }
        
        answer_slice = append(answer_slice, min_num)
        
    }
    return answer_slice
}