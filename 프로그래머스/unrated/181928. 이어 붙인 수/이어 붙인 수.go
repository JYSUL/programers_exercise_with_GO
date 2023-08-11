func solution(num_list []int) int {
    var odd_sum, even_sum int = 0, 0
    for _, num := range num_list {
        if num % 2 == 0 {
            even_sum = even_sum * 10 + num
        } else {
            odd_sum = odd_sum * 10 + num
        }
    }
    return odd_sum + even_sum
}