import "sort"
func solution(arr []int, divisor int) []int {
    var result []int
    for _, num := range arr {
        if num % divisor == 0 {
            result = append(result, num)
        }
    }
    if len(result) == 0 {
        return []int{-1}
    } else {
        sort.Slice(result, func(i,j int)bool {
            return result[i] < result[j]
        })
    }
    
    return result
}