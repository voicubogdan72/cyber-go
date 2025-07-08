package main

import (
	"bufio"
	"fmt"
	"os"
)


var key byte = 0xAA

func xorEncryptDecrypt(input string) string{
	ouput := make([]byte, len(input))
	for i := range input{
		ouput[i] = input[i] ^ key
	}
	return  string(ouput)
}

func main(){
	fmt.Println("Keylogger active")

	file, err := os.OpenFile("log.enc", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		fmt.Println("Eroare la deschiderea fișierului:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		encrypted := xorEncryptDecrypt(line)
		_, err := file.WriteString(encrypted + "\n")
		if err != nil {
			fmt.Println("Eroare la scrierea fișierului:", err)
			break
		}
	}
}
