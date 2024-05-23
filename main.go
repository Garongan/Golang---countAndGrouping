package main

import (
	"bufio"
	countWord "com/enigma/countAndSort/countAndSort/countWord"
	groupWord "com/enigma/countAndSort/countAndSort/groupingWord"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Masukkan kalimat: We Always Mekar")
	input, _ := reader.ReadString('\n')
	fmt.Println(countWord.CountWord(strings.TrimSpace(input)))

	fmt.Println("Masukkan array kata (dipisahkan spasi): Pendanaan Terproteksi Untuk Dampak Berarti")
	inputStrings, _ := reader.ReadString('\n')
	fmt.Printf("%v\n", groupWord.GroupWord(strings.TrimSpace(inputStrings)))
}
