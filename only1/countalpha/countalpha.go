package exam

// import "fmt"

func CountAlpha(s string) int {
	count := 0
	for i := 0; i < len(s); i++ {
		if (s[i] >= 'a' && s[i] <= 'z') {
			count++
		}
		if (s[i] >= 'A' && s[i] <= 'Z') {
			count++
		}
	}
	return count
}

// func main() {
// 	s1 := "hello hello 99 //la"
// 	// s2 := "HELLO world"
// 	// s3 := "H e LL o 69 ..."
// 	res1 := CountAlpha(s1)
// 	// res2 := CountAlpha(s2)
// 	// res3 := CountAlpha(s3)

// 	fmt.Println(res1)
// 	// fmt.Print('\n')
// 	// fmt.Println(res2)
// 	// fmt.Print('\n')
// 	// fmt.Println(res3)
// 	// fmt.Print('\n')
//}