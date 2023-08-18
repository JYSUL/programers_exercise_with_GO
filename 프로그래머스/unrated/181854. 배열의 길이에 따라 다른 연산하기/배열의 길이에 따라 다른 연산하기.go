func solution(arr []int, n int) []int {
    check := len(arr) % 2
    for i := 0; i < len(arr); i++ {
        if i % 2 != check {
            arr[i] += n
        }
    }
    return arr
}