import "strings"

func babbling_check (bab string) int {
    var babbles []string = []string{"aya", "ye", "woo", "ma"}
    var doubles []string = []string{"ayaaya", "yeye", "woowoo", "mama"}
    for _, double := range doubles {
        if strings.Contains(bab, double) {
            return 0
        }
    }
    for _, babble := range babbles {
        bab = strings.Replace(bab, babble, " ", -1)
    }
    if len(strings.Replace(bab, " ","", -1)) == 0{
        return 1
    }
    return 0
}


func solution(babbling []string) int {
    var answer int
    for _, bab := range babbling {
        answer += babbling_check(bab)
    }
    
    return answer
}