func solution(number []int) int {
    var answer int = 0
    for i := 0; i < len(number); i++ {
        for j := i+1; j < len(number); j++ {
            for k := j+1; k < len(number); k++ {
                if number[i] + number[j] + number[k] == 0 {
                    answer += 1
                }
            }
        }
    }
    return answer
}