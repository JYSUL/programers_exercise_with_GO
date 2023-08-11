func solution(arr []int, queries [][]int) []int {
    for _, query := range queries {
        arr[query[0]], arr[query[1]] = arr[query[1]], arr[query[0]]
    }
    return arr
}