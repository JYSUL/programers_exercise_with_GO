//import "fmt"
func solution(arr []int) int {
    var flag bool = true
    var count int = 0
    for flag {
        flag = false
        for index, num := range arr {
            if num >= 50 && num % 2 == 0 {
                arr[index] /= 2
                flag = true
            } else if num < 50 && num % 2 == 1 {
                arr[index] *= 2 + 1
                flag = true
            }
        }
        if count == 0 && flag == false {
            return 0
        }
        count += 1
    }
    return count
}