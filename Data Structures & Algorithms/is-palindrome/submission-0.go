func isPalindrome(s string) bool {
    s = strings.ToLower(s)
    reg := regexp.MustCompile(`[^a-z0-9]+`)
	s = reg.ReplaceAllString(s, "")
    fmt.Println(s)
    runes := []rune(s)
    j := len(runes) - 1
    for i := 0; i < len(runes); i ++ {
        if runes[i] != runes[j] {
            return false
        }
        j--
    }
    return true
}
