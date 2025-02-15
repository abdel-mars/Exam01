package exam

// import "fmt"

func PrintIf(str string) string {
	res := "G"
	if str == "" {
		return res
	} else if len(str) >= 3 {
		return res
	} else {
		res = "Invalid Input"
		return res
	}
}

// func main() {
// 	fmt.Println(PrintIf("abcdefz"))
// 	fmt.Println(PrintIf("abc"))
// 	fmt.Println(PrintIf(""))
// 	fmt.Println(PrintIf("14"))
// }
