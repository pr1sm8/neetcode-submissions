func groupAnagrams(strs []string) [][]string {
	count := make(map[[26]int][]string)
	for _, str := range strs {
		hash := [26]int{}
		for _, char := range str {
			hash[char-'a'] ++
		}
		count[hash] = append(count[hash], str)
	}
	res := make([][]string, len(count))
	i := 0
	for _, v := range count {
		res[i] = v
		i++
	}
	return res
}
