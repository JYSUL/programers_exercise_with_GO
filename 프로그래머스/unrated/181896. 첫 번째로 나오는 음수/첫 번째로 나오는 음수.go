func solution(num_list []int) int {
    for index, num := range num_list {
        if num < 0 {
            return index
        }
    }
    return -1
}