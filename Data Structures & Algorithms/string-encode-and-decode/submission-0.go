type Solution struct{}

func (s *Solution) Encode(strs []string) string {
    var sb strings.Builder
    for _, str := range strs {
        sb.WriteString(fmt.Sprintf("%dé%s", len(str), str))
    }
    return sb.String()
}

func (s *Solution) Decode(encoded string) []string {
    var res []string
    var encodedR = []rune(encoded)
    // for i, r := range encodedR {
    //     fmt.Printf("r: %c, i: %d\n", r, i)
    // }
    var lenC []rune
    for i := 0; i < len(encodedR); i++ {
        r := encodedR[i]
        // fmt.Printf("rune at i: %d, r: %c\n",i,r)
        
        if r == 'é' {
            // fmt.Println("lenC", string(lenC))
            lenI, err := strconv.Atoi(string(lenC))
            if err !=nil {
                // fmt.Println("error in atoi", err)
                return []string{}
            }
            lenC = []rune{}
            start := i+1
            end := i+1+lenI
            // fmt.Printf("start: %d, end: %d, slice: %s\n", start, end, string(encodedR[start:end]))
            res = append(res, string(encodedR[start:end]))
            i = i + lenI
        } else {
            lenC = append(lenC, r)
        }
    }
    return res
}
