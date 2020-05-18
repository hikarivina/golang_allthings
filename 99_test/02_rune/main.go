package main

// https://qiita.com/seihmd/items/4a878e7fa340d7963fee
import "fmt"

func main() {
	// s := "abcde"

	// for i := 0; i < len(s); i++ {
	// 	b := s[i]      // byte
	// 	fmt.Println(b) // 227, 129, 130...
	// }

	// for _, r := range s {
	// 	// rune
	// 	fmt.Println(r) // 12354, 12356, 12358
	// }

	s1 := "あ"
	for i := 0; i < len(s1); i++ {
		fmt.Printf("% x", s1[i]) // e3 81 82
	}
}
