func solution(arr []int, query []int) []int {
    //var odd_seq, even_seq int = 0, 0
    for index, num := range query {
        if index % 2 == 0 {
            arr = arr[:num+1]
        } else {
            arr = arr[num:]
        }
    }
    return arr
}