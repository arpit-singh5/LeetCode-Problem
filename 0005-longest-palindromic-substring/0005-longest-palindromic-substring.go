func longestPalindrome(s string) string {
    maxStart := 0
    maxLen := 1
    for i := 0; i < len(s); i++ {
        for j := i+1; j <= len(s); j++ {
            substring := s[i:j]
           
            if isPalindrome(substring) && len(substring) > maxLen {
                maxLen = len(substring)
                maxStart = i
            }
        }
    }
    return s[maxStart:maxLen+maxStart]
}

func isPalindrome(s string) bool {
    left, right := 0, len(s)-1
    for left < right {
    if s[left] != s[right] {
        return false
    }
    left++
    right--
    }
    return true
}