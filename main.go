package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

func main() {
	// Start
	println("コンソールタイピングゲームへようこそ")
	print("名前を入力してください:")
	var name string
	fmt.Scan(&name)

	// Game
	words := []string{"neko", "ablaze", "golang", "linux", "sushi", "bento", "syaneko", "tsuihai", "game", "code", "aqua", "sudo", "apt", "happy", "floorp", "334"}
	var socre int = 0

	println("次の文字を入力してください\n")
	for i := 0; i < 10; i++ {
		num := rand.Intn(len(words))
		fmt.Printf("[%s]\n", words[num])
		var input string
		fmt.Scan(&input)
		if input == words[num] {
			println("- OK\n")
			socre++
		} else {
			println("- NG\n")
		}
	}

	// Result
	fmt.Printf("%sさんのスコアは %d/10 です\n終了するにはCtrl + C又は閉じてください", name, socre)
	var sc = bufio.NewScanner(os.Stdin)
	for sc.Scan() {
	}
}
