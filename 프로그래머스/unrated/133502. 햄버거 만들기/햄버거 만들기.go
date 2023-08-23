//import "fmt"
//import "strconv"


func solution(ingredients []int) int {
    var answer int
    var Hamburger []int = []int{1, 2, 3, 1}
    var check []int
    for _, ingredient := range ingredients {
        check = append(check, ingredient)
        flag := false
        if len(check) > 3 {
            flag = true
            for i := 0; i < 4; i++ {
                if check[len(check)-4+i] != Hamburger[i] {
                    flag = false
                    break
                }
            }
        }
        if flag {
            answer++
            check = check[:len(check)-4]
        }
        
        
    }
    
    return answer
}

/*
빵1 야채2 고기3 빵 1
*/