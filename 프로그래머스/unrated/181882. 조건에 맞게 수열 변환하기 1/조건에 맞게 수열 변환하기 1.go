func solution(arr []int) []int {
    for index, num := range arr {
        if num >= 50 && num % 2 == 0 {
            arr[index] /= 2
        } else if num < 50 && num % 2 == 1 {
            arr[index] *= 2
        }
    }
    return arr
}