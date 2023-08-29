import (
    //"fmt"
    //"math"
)
func solution(n int) int64 {
    var priv, now int64 = 1, 1
    for i := 0; i < n-1; i++ {
        priv, now = (priv + now) %1234567, priv
    }
    return priv
}