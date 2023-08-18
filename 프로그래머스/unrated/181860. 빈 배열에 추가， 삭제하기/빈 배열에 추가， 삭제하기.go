//import "fmt"
func solution(arr []int, flag []bool) []int {
    var answer []int
    
    for index, spin := range flag {
        if spin {
            for i := 0; i < arr[index] * 2; i++ {
                answer = append(answer, arr[index])
            }
        } else {
            answer = answer[:len(answer)-arr[index]]
        }
        //fmt.Println(answer)
    }   
    
    return answer
}