func solution(strArr []string) []string {
    var answer_slice []string
    var flag bool
    for _, str := range strArr {
        flag = true
        for i := 0; i < len(str) - 1; i++ {
            if str[i:i+2] == "ad" {
                flag = false
                break
            }
        }
        if flag {
            answer_slice = append(answer_slice, str)
        }
    }
    
    return answer_slice
}