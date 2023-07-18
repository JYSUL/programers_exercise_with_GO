func solution(num_list []int) [2]int {
    var even int = 0
    var odd int = 0
    for _, num := range num_list {
        if num % 2 == 0 {
            even++
        } else {
            odd++
        }
    }
    return [2]int{even,odd}
}