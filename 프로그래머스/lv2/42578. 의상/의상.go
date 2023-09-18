//import "fmt"


func solution(clothes [][]string) int {
    var answer int = 1
    var select_check map[string]map[string]bool = make(map[string]map[string]bool)  //중복확인을 위한 해시테이블
    var select_count map[string]int = make(map[string]int)                          //확인용 해시테이블
    
    //fmt.Println(select_count)
    //fmt.Println(select_check)
    
    for _, cloth := range clothes {
        check, selection := cloth[0], cloth[1]
        //map in map을 위한 초기화
        if _, flag := select_check[selection]; !flag {  
            select_check[selection] = make(map[string]bool)
        }
        
        //중복이 아니면 종류의 count++
        if !select_check[selection][check] {
            select_check[selection][check] = true
            select_count[selection]++
        }
    }
    
    //fmt.Println(select_count)
    
    for _, value := range select_count {
        answer *= value + 1
    }
    
    //모두 벗은 경우의 수 제거
    return answer - 1
}