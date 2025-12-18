# Group Anagrams

```go
package main

import "fmt"

func sort(str string) string {
	chars := []rune(str)

	for i := 1; i < len(chars); i++ {
		key := chars[i]
		j := i - 1

		for j >= 0 && int(chars[j]) > int(key) {
			chars[j + 1] = chars[j]
			j--
		}

		chars[j + 1] = key
	}

	return string(chars)
}

func main() {
	strs := []string{"act", "pots", "tops", "cat", "stop", "hat"}

	m := make(map[string][]string)
	res := [][]string{}

	for _, v := range strs {
		sorted := sort(v)

		if _, ok := m[sorted]; ok {
			m[sorted] = append(m[sorted], v)
		} else {
			m[sorted] = []string{v}
		}
	}

	for _, v := range m {
	    res = append(res, v)
    }

	fmt.Printf("%v", res)
}

```
