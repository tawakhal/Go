package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	nama string
	umur int
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Siapa nama anda ? ")
	scanner.Scan()
	nama = scanner.Text()

	fmt.Println("Siapa nama panjang anda ? ")
}
