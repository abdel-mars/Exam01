package main

// import "fmt"

func FirstWord(s string) string {
	
	res := ""

	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			break
		}
		res += string(s[i])
	}
	return res + "\n"
}

// func main() {
//     fmt.Print(FirstWord("hello there"))
//     fmt.Print(FirstWord(""))
//     fmt.Print(FirstWord("hello   .........  bye"))
// }
