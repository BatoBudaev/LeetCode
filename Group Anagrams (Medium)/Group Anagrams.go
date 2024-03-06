package main

import "fmt"

func groupAnagrams(strs []string) [][]string {
	m := make(map[[26]rune][]string)

	for _, s := range strs {
		k := [26]rune{}

		for i := range s {
			k[s[i]-'a'] += 1
		}

		m[k] = append(m[k], s)
	}

	var res [][]string

	for _, v := range m {
		res = append(res, v)
	}

	return res
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	groups := groupAnagrams(strs)
	fmt.Println(groups)
}
