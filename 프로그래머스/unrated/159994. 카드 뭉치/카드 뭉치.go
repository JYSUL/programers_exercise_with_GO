func solution(cards1 []string, cards2 []string, goal []string) string {
    var index_cards1, index_cards2 int
    for index_goal := 0; index_goal < len(goal); index_goal++ {
        if index_cards1 < len(cards1) && cards1[index_cards1] == goal[index_goal] {
            index_cards1++
        } else if index_cards2 < len(cards2) && cards2[index_cards2] == goal[index_goal] {
            index_cards2++
        } else {
            return "No"
        }
    }
    return "Yes"
}