func lengthOfLongestSubstring(s string) int {
    maxLength := 0
   charLastSeen := make(map[rune]int)
   last := 0
   for right := 0; right < len(s); right++ {
    c := rune(s[right])
    if index, ok := charLastSeen[c]; ok {
        if index >= last{
            last = index +1
        }
    }
    charLastSeen[c] = right
    currentWindow := right-last+1
    if currentWindow > maxLength {
        maxLength = currentWindow
    }

   }
   return maxLength
}