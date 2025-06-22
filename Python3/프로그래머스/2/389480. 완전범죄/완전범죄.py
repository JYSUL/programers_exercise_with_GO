def solution(info, n, m):
    total_A = 0
    total_B = 0
    info.sort(key = lambda x : (x[0] / x[1]), reverse = True) 
    for x0, x1 in info :
        total_A += x0
    
    for x0, x1 in info :
        if total_B + x1 < m :
            total_A -= x0
            total_B += x1
    
    if total_A >= n :
        return -1
    else :
        return total_A
    


"""
대락적으로만 가늠해서 짠 알고리즘
반례가 존재하지 않는다는 확신은 없다
"""