func solution(n int, left int64, right int64) []int {
    var answer []int
    for i := left; i <= right; i++ {
        if i % int64(n) > i / int64(n) {  //열이 행보다 클때
            answer = append(answer, int(i % int64(n) + 1))
        } else {
            answer = append(answer , int(i / int64(n) + 1))
        }
    }
    return answer
}


/*
123456789
123223333
*/