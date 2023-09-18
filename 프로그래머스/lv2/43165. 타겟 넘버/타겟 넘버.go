func solution(numbers []int, target int) int {
    if len(numbers) == 0 {
        if target == 0 {
            return 1
        } else {
            return 0
        }
    } else {
        return solution(numbers[1:],target+numbers[0]) + solution(numbers[1:], target-numbers[0])
    }
}
/*
재귀함수

맨압이 +일때와 -일때를 나누어서
top down한다.


*/