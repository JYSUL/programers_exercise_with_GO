func solution(ineq string, eq string, n int, m int) int {
    var flag int = 0
    
    if eq == "=" {
        if n == m {
            flag = 1
        }
    }
    
    if ineq == "<" {
        if n < m {
            flag = 1
        }
    } else {
        if n > m {
                flag = 1
        }
    }    
    return flag
}