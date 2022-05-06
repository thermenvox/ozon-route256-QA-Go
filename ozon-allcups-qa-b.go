// https://cups.online/ru/tasks/1250

package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var actions []string
	var reverseIndices []int
	reverseIndices = append(reverseIndices, 0)
	for i := 0; i < n; i++ {
		var mod string
		fmt.Scan(&mod)
		if mod == "!" {
			reverseIndices = append(reverseIndices, i)
		}
		actions = append(actions, mod)
	}
	reverseIndices = append(reverseIndices, len(actions)-1)

	var str string = ""
	for i := 1; i < len(reverseIndices); i++ {
		if len(reverseIndices[i:])%2 != 0 {
			for j := reverseIndices[i-1]; j < reverseIndices[i]; j++ {
				if actions[j] != "!" {
					str = str + string(actions[j][1])
				}
			}
		} else {
			for j := reverseIndices[i-1]; j < reverseIndices[i]; j++ {
				if actions[j] != "!" {
					str = string(actions[j][1]) + str
				}
			}
		}
	}

	for i := reverseIndices[len(reverseIndices)-1]; i < len(actions); i++ {
		if actions[i] != "!" {
			str = str + string(actions[i][1])
		}
	}
	fmt.Println(str)
}
