func solution(arr []int, idx int) int {
    var answer int = -1
    for i := idx; i < len(arr); i++ {
        if arr[i] == 1 {
            answer = i
            break
        }
    }
    return answer
}