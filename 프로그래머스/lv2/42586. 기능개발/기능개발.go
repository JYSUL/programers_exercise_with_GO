//import "fmt"
func solution(progresses []int, speeds []int) []int {
    var check_list []int
    var result []int
    var check_table map[int]int = make(map[int]int)
    
    //걸리는 날짜가 내 앞의 것보다 작으면 내 앞의 것으로 간주하고 해당 날짜의 끝나는 작업 count++
    
    //맨처음을 위한 사전작업
    pre_check := (100-progresses[0]) / speeds[0]
        if (100-progresses[0]) % speeds[0] != 0 {
            pre_check++
    }
    check_table[pre_check]++
    check_list = append(check_list, pre_check)
    
    //걸리는 날짜가 바뀌면 check_list에 append -> 나중에 순서대로 check_table을 순회하기 위해서 -> key로 리스트 만들어서 sort해도 되긴한다.
    //작업이 끝나는 날짜의 count++
    for i := 1; i < len(progresses); i++ {
        check := (100-progresses[i]) / speeds[i]
        if (100-progresses[i]) % speeds[i] != 0 {
            check++
        }
        if check > pre_check {
            pre_check = check
            check_list = append(check_list, pre_check)
        }
        check_table[pre_check]++
    }
    
    //fmt.Println(check_list)
    //fmt.Println(check_table)
    
    for _, check := range check_list {
        result = append(result, check_table[check])
    }
    
    return result
}