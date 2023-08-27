def solution(s):
    sList = s.split(" ")
    answerList = ["" for i in range(len(sList))]
    for i in range(len(sList)) :
        for j in range(len(sList[i])) :
            if j % 2 == 0 :
                answerList[i] += sList[i][j].upper()
            else :
                answerList[i] += sList[i][j].lower()
    print(answerList)
    return " ".join(answerList)