func solution(num_list []int) int {
    var even_index_num, odd_index_num int = 0, 0
    for index, num := range num_list {
        if index % 2 == 0 {
            even_index_num += num
        } else {
            odd_index_num += num
        }
    }
    
    if even_index_num > odd_index_num {
        return even_index_num
    } else {
        return odd_index_num
    }
    
}