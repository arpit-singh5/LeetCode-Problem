func mergeAlternately(word1 string, word2 string) string {
    var mergeWord string
    n := max(len(word1), len(word2))
    for i := 0; i< n; i++ {
        if i < len(word1) {
            mergeWord += string(word1[i])
        }
        if i < len(word2) {
            mergeWord += string(word2[i])
        }
    }
    return mergeWord
}

func max(a,b int) int {
    if a > b {
        return a
    }
    return b
}