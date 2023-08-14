func solution(num_list []int) int {
    var answer, count int = 0, 0
    for _, num := range num_list {
        count = 0
        for num != 1 {
            if num % 2 == 0 {
                num /= 2
            } else {
                num = (num - 1) / 2
            }
            count += 1
        }
        answer += count
    }
    return answer
}