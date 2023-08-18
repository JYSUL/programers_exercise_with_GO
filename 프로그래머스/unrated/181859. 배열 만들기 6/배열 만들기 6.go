//import "fmt"
func solution(arr []int) []int {
    var answer []int
    for index, num := range arr {
        if len(answer) == 0 {
            answer = append(answer,num)
            
        } else if answer[len(answer)-1] == arr[index] {
            answer = answer[:len(answer)-1]
            
        } else {
            answer = append(answer,num)
        }
        //fmt.Println(answer)
    }
    if len(answer) == 0 {
        answer = []int{-1}
    }
    return answer
}