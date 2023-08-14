func solution(arr []int) []int {
    var flag bool = false
    var start, end int = 0, 0
    for i := 0; i < len(arr); i++ {
        if arr[i] == 2 {
            start = i
            flag = true
            break
        }
    }
    if !flag {
        return []int{-1}
    }
    
    for i := len(arr) -1; i >= 0; i-- {
        if arr[i] == 2 {
            end = i
            break
        }
    }
    
    return arr[start: end+1]
}