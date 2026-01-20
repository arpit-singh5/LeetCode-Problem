func gcdOfStrings(str1 string, str2 string) string {
    if (str1+str2) != (str2+str1) {
        return ""
    }
    cd := gcd(len(str1),len(str2))
    return str1[0:cd]

}
func gcd(a,b int) int {
   for b != 0 {
    temp := b
    b = a % b
    a = temp
   }
   return a
}