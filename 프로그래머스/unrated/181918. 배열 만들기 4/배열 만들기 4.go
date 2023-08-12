func solution(arr []int) []int {
    var i int = 0
    var answer_slice []int
    for i < len(arr) {
        if len(answer_slice) == 0 {
            answer_slice = append(answer_slice, arr[i])
            i++
        } else {
            if arr[i] > answer_slice[len(answer_slice)-1] {
                answer_slice = append(answer_slice, arr[i])
                i++
            } else {
                answer_slice = answer_slice[:len(answer_slice)-1]
            }
        }
    }
    return answer_slice
}