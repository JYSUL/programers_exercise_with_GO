func sumIntSlice(elements []int) int {
    var sum int
    for i := 0; i < len(elements); i++ {
        sum += elements[i]
    }
    return sum
}
func solution(elements []int) int {
    check_list := map[int]bool{}
    elements = append(elements, elements...)
    for i := 0; i < len(elements) / 2; i++ {
        for j := 0; j < len(elements) / 2; j ++ {
            check_list[sumIntSlice(elements[i:i+j])] = true
        }
    }
    return len(check_list)
}