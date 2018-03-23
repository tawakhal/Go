package main

import "fmt"

func main() {
	messages := make(chan string) /*pendeklarasian pembuatan channel dengan nama
	  messages yang bertipe string*/

	go func() {
		messages <- "ping" /*membuat function go yang dimana di deklarasikan
		  dengan nama messages yang berisi comment "ping" yang akan di kirim ke channel
		  messages*/

	}()
	msg := <-messages //membuat variabel msg untuk memanggil channel messages
	fmt.Println(msg)
}
