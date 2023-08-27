func solution(s string) string {
    return s[(len(s)-1)/2 : len(s)/2 +1]
}

/*
홀수 ex 5 -> [2:3]
짝수 ex 4 -> [1:3]

*/