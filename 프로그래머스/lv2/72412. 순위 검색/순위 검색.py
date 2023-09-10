import bisect
def solution(info, query):
    answers = []
    info  = list(map(lambda x: x.split()                    ,info))
    query = list(map(lambda x: x.replace(" and","").split() ,query))
    info_table = dict()
    
    for i in ["java","python","cpp","-"] :
        for j in ["backend", "frontend","-"] :
            for k in ["senior","junior","-"] :
                for l in ["pizza","chicken","-"] :
                    info_table[i+j+k+l] = []
    for language, job, career, food, score in info :
        score = int(score)
        for i in [language,"-"] :
            for j in [job,"-"] :
                for k in [career,"-"] :
                    for l in [food,"-"] :
                        info_table[i+j+k+l].append(score)
    for scores in info_table.values() :
        scores.sort()
    
    for language, job, career, food, score in query :
        score = int(score)
        checking_list = info_table[language+job+career+food]
        """
        start, end = 0, len(checking_list)-1        
        while start <= end : # binary search
            mid = (start + end) // 2
            if checking_list[mid] < score :
                start = mid + 1
            else :
                end = mid - 1
        """
        answers.append(len(checking_list) - bisect.bisect_left(checking_list, score))
    
    
    
    
    return answers



"""
python 으로 데이터베이스 구현?
term으로 걸르고 걸르고 하지말고 info 하나한 확인하는게 나은가?

"""