from collections import deque
def solution(order):
    m = len(order)
    order = deque(order)
    stacks = deque()
    boxs = deque([i for i in range(1,len(order)+1)])
    while order :
        if boxs :
            if boxs[0] == order[0] :
                boxs.popleft()
                order.popleft()
                continue
        if stacks :
            if stacks[-1] == order[0] :
                stacks.pop()
                order.popleft()
                continue
        
        if boxs :
            stacks.append(boxs.popleft())
        else :
            break
        
        
    return m - len(order)