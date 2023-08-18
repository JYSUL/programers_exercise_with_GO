func solution(arr []int) []int {
    var bina int = 1
    for bina < len(arr) {
        bina *= 2
    }
    for i := len(arr); i < bina; i++ {
        arr = append(arr, 0)
    }
    return arr
}