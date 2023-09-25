def solution(numbers):
    answer = []
    for i in numbers :
        binum = bin(i)[2:]
        if "0" in binum :
            if binum[-1] == "0" :
                answer.append(int(binum[:-1]+"1", 2))
            else :
                answer.append(int(binum[::-1].replace("10", "01", 1)[::-1], 2))
                        
        else :
            answer.append(int("10"+binum[1:], 2))
    return answer