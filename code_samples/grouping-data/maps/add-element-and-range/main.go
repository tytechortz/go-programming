package main

import "fmt"

func main() {
	m := map[string]int{
		"James":32,
		"Miss Moneypenny":27,
	}
	fmt.Println(m)
	fmt.Println(m["James"])

	v, ok := m["Barnabas"]
	fmt.Println(v)
	fmt.Println(ok)

	m["Todd"] = 33

	if v, ok := m["James"]; ok {
		fmt.Println("IF PRINT", v)
	}

	for k, v := range m {
		fmt.Println(k, v)
	}

}