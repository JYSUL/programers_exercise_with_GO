import (
    //"fmt"
)
func solution(arr1 [][]int, arr2 [][]int) [][]int {
    if len(arr1[0]) != len(arr2) {
        panic("Can't product given matrixes")
    }
    var table [][]int
    
    //fmt.Println(table)
    //출력 길이의 맞게 출력할 행렬의 요소를 0으로 초기화
    for i := 0; i < len(arr1); i++ {
        table = append(table, []int{})
        for j := 0; j < len(arr2[0]); j++ {
            table[i] = append(table[i], 0)
        }
    }
    
    //출력할 행렬을 순회하면서 값의 계산
    for i := 0; i < len(table); i++ {
        for j := 0; j < len(table[0]); j++ {
            for k := 0; k < len(arr1[0]); k++ {
                table[i][j] += arr1[i][k] * arr2[k][j]
            }
        }
    }
    
    return table
}


