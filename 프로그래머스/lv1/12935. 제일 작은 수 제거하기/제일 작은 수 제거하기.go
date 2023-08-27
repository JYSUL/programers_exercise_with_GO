func solution(arr []int) []int {
    if len(arr) == 1 {
        return []int{-1}
    }
    var min_index, min_num int = 0, arr[0]
    for index, num := range arr {
        if num < min_num {
            min_index, min_num = index, arr[index]
        }
    }
    return append(arr[:min_index], arr[min_index+1:]...)
}