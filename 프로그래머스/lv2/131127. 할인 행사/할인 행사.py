def solution(want, number, discount):
    answer = 0
    for i in range(len(discount)-10 + 1) :
        discount_check = discount[i:i+10]
        for i, j in zip(want, number) :
            if discount_check.count(i) != j :
                answer -= 1
                break
        answer += 1
            
    return answer