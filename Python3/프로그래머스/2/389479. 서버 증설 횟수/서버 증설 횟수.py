def solution(players, m, k):
    answer = 0
    available_server = m
    #k시간 전부터 반납 전까지 증설한 서버내역을 위한 list
    servers = []
    for i in range(k) :
        servers.append(0)
        
    
    for i in range(len(players)) :
        #k시간 전에 증설한 서버의 반환
        server_count = i%k - k
        available_server -= servers[server_count] * m
        servers[server_count] = 0
        
        #서버가 부족하면 부족한 만큼 추가로 증설 / if문과 //연산이 좀더 나을 수도?
        while available_server <= players[i] :
            available_server += m
            servers[i%k] += 1
            answer += 1
            
        #print(servers)
        #print(available_server)
        
    return answer