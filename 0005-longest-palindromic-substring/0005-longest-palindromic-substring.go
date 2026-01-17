func longestPalindrome(s string) string {
    if len(s) == 0 {
        return ""
    }
    C := 0
    R := 0  
    mString := ""

    mString += "#"
    for _,v := range s {
        mString += string(v) + "#"
    }
    n := len(mString)
    p := make([]int, n)

    for i,_ := range mString {
        mirror := 2*C - i

        if i < R {
            p[i] = mini(R-i, p[mirror])
        }
        for i+p[i]+1 < n && i-p[i]-1 >= 0 && mString[i+p[i]+1] == mString[i-p[i]-1] {
            p[i]++
        }
        if i+p[i] > R {
            C = i
            R = i+ p[i]
        }
    }
    maxLen := 0
    Center := 0
    for i := 1; i < n-1; i++ {
        if p[i] > maxLen {
            maxLen = p[i]
            Center = i
        }
    }
    start := (Center - maxLen)/2
    return s[start : start+maxLen]
}
func mini (a,b int) int {
    if a>b {
        return b
    }
    return a
}