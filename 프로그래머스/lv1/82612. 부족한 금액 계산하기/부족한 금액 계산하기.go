func solution(price int, money int, count int) int64 {
    var total_price int = 0
    
    for i := 0; i <= count; i++ {
        total_price += price * i
    }
    
    if total_price <= money {
        return 0
    } else {
        return int64(total_price - money)
    }
}