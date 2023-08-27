//import "fmt"
func solution(arr1 [][]int, arr2 [][]int) [][]int {
    for i := 0; i < len(arr1); i++ {
        for j := 0; j < len(arr1[0]); j++ {
            arr1[i][j] += arr2[i][j]
        }
    }
    return arr1
}