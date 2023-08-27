func solution(price int, money int, count int) int64 {
    var total_price int = price * count * (count+1) / 2
    
    
    if total_price <= money {
        return 0
    } else {
        return int64(total_price - money)
    }
}