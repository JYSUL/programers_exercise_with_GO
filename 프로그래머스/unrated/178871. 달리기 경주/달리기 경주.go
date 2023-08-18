//import "fmt"
func solution(players []string, callings []string) []string {
    seq_table := make(map[string]int)
    for index, player := range players {
        seq_table[player] = index
    }
    
    for _, calling := range callings {
        rate := seq_table[calling]
        seq_table[calling]--
        seq_table[players[rate-1]]++
        players[rate], players[rate-1] = players[rate-1], players[rate]
    }
    return players
}