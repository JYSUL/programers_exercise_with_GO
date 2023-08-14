def solution(arr, query):
    arr_check = [] + arr
    oddnum, evennum = 0, 0
    for i, x in enumerate(query) :
        if i% 2 :
            arr = arr[x:]
        else :
            arr = arr[:x+1] 
            
    return arr

#print(solution(list(range(20)), [10,-2,9,4,8,0,8,0,10,0]))


"""
왜 이 코드는 안됨?
if i% 2 :
    arr = arr[x:]
else :
    arr = arr[:x] 
이게 정답이 맞고 해당 문제가 오류로 보임
채점 테스트에서 마지막 제한사항 query의 각 원소는 0보다 크거나 같고 남아있는 arr의 길이 보다 작습니다 를 위반하는 테스트가 있는거로 추정
반례)
[0,1,...,19] 에 query가 [10,0,20,0,20]는 맨앞 10개를 자르고 0~10에서만 처리되야하는게
잘못된 정답이 나오는 코드가 테스트들에서는 정답처리가 되고있음

3,10 번은 empty의 경우의 수
쿼리의 원소는 0이상
"""