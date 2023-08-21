import (
    "container/heap"
    //"fmt"
)
type IntHeap []int
 
func (h IntHeap) Len() int {
    return len(h)
}
 
func (h IntHeap) Less(i, j int) bool {
    return h[i] < h[j]
}
 
func (h IntHeap) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}
 
func (h *IntHeap) Push(element interface{}) {
    *h = append(*h, element.(int))
}
 
func (h *IntHeap) Pop() interface{} {
    old := *h
    n := len(old)
    element := old[n-1]
    *h = old[0 : n-1]
    return element
}


func solution(k int, scores []int) []int {
    var answer_slice []int
    h := &IntHeap{}
    for _, score := range scores {
        heap.Push(h, score)
        if h.Len() > k {
            heap.Pop(h)
        }
        now_heap := *h
        answer_slice = append(answer_slice, now_heap[0])
    }
    return answer_slice
}