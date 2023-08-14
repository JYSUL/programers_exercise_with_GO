func solution(num_list []int) int {
    var answer int 
    if len(num_list) > 10 {
        answer = 0
        for _, num := range num_list {
            answer += num
        } 
    } else {
        answer = 1
        for _, num := range num_list {
            answer *= num
        }
    }
    return answer
}