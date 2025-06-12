def solution(s):
    list_string = []
    for num in s.split() :
        list_string.append(int(num))
    
    min_num, max_num = list_string[0], list_string[0]
    for num in list_string :
        if min_num > num : min_num = num
        if max_num < num : max_num = num
    
    
    return str(min_num) + " " + str(max_num)