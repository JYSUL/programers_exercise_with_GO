import (
    //"fmt"
    //"combin"
)

func isPrime(num int) bool {
    flag := true
    for i := 2; i <= num / 2; i++ {
        if num % i == 0 {
            flag = false
            break
        }
    } 
    
    return flag
    
}


func solution(nums []int) int {
    var answer int = 0
    for i := 0; i < len(nums); i++ {
        for j := i+1; j < len(nums); j++ {
            for k := j+1; k < len(nums); k++ {
                //fmt.Println(nums[i] + nums[j] + nums[k])
                if isPrime(nums[i] + nums[j] + nums[k]) {
                    answer++
                }
            }
        }
    }
    
    
    return answer

}