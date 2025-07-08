package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


func main(){
	fmt.Print("Put the password: ")
	var inputPassword string
	fmt.Scanln(&inputPassword)


	file, err := os.Open("wordlist.txt")
	if err != nil{
		fmt.Println("Error open the file: ", err)

		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	found := false

	for scanner.Scan(){
		word := strings.TrimSpace(scanner.Text())
		if word == inputPassword{
			found = true
			break
		}
	}

	if err := scanner.Err(); err != nil{
		fmt.Println("Error read the file: ", err)
		return
	}

	if found{
		fmt.Println("Password was found! Is not safe")
	}else{
		fmt.Println("Password was nit found! Is safe")

	}
}