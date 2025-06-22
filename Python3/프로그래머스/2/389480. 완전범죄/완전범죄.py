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
    
    return 0