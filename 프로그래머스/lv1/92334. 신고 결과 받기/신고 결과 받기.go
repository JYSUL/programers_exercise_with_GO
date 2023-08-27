import (
    "strings"
    //"fmt"
)
func solution(id_list []string, reports []string, k int) []int {
    reported_table := make(map[string][]string)
    reporter_count := make(map[string]int)
    true_reports := []string{}
    true_reports_check := make(map[string]bool)
    var answer []int
    
    //reports 에서 중복제거를 한 새 reports를 만든다.
    for _, report := range reports {
        if !true_reports_check[report] {
            true_reports_check[report] = true
            true_reports = append(true_reports, report)
        }
    }
    
    
    
    //신고 당한사람을 기준으로 그사람을 신고한 목록을 map 형태로 정리
    for _, report := range true_reports {
        reporter := strings.Split(report, " ")[0]
        reported := strings.Split(report, " ")[1]
        
        reported_table[reported] = append(reported_table[reported], reporter)
    }
    
    //신고당한 사람을 신고한 총 사람수가 k이상일때, 신고한 사람의 count++
    for _, id := range id_list {
        if len(reported_table[id]) >= k {
            for _, reporter := range reported_table[id] {
                reporter_count[reporter]++
            }
        }
    }
    for _, id := range id_list {
        answer = append(answer, reporter_count[id])
    }
    /*
    fmt.Println(reported_table)
    fmt.Println(reporter_count)
    fmt.Println(answer)
    */
    return answer
}