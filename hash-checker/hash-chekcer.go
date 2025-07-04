package main

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"os"
)



func hashString(s string)(string, string){
	md5Hash := md5.Sum([]byte(s))

	sha256Hash := sha256.Sum256([]byte(s))

	return hex.EncodeToString(md5Hash[:]), hex.EncodeToString(sha256Hash[:])
}

func hashFile(path string)(string, string, error){
	file, err := os.Open(path)
	if err != nil {
		return "", "", err
	}

	defer file.Close()

	md5Hash := md5.New()
	sha256Hash := sha256.New()
	
	_, err = io.Copy(io.MultiWriter(md5Hash,sha256Hash), file)
	if err != nil{
		return "","", err
	}

	return hex.EncodeToString(md5Hash.Sum(nil)), hex.EncodeToString(sha256Hash.Sum(nil)), nil
}


func main(){
	// Flags pentru modul de input
	textPtr := flag.String("text", "", "String text pentru hash")
	filePtr := flag.String("file", "", "Cale către fișier pentru hash")
	checkPtr := flag.String("check", "", "Hash pentru verificare")
	algoPtr := flag.String("algo", "sha256", "Algoritm de hash pentru verificare: md5 sau sha256")

	flag.Parse()


	var md5Hash, sha256Hash string
	var err error

	if *textPtr == "" && *filePtr == "" {
		fmt.Println("Trebuie să specificați fie -text, fie -file")
		return
	}

	if *textPtr != "" {
		md5Hash, sha256Hash = hashString(*textPtr)
	} else if *filePtr != "" {
		md5Hash, sha256Hash, err = hashFile(*filePtr)
		if err != nil {
			fmt.Printf("Eroare la citirea fișierului: %v\n", err)
			return
		}
	}

	fmt.Printf("MD5:    %s\n", md5Hash)
	fmt.Printf("SHA256: %s\n", sha256Hash)

	// Dacă s-a cerut verificare hash
	if *checkPtr != "" {
		match := false
		switch *algoPtr {
		case "md5":
			match = (*checkPtr == md5Hash)
		case "sha256":
			match = (*checkPtr == sha256Hash)
		default:
			fmt.Println("Algoritm necunoscut pentru verificare. Folosiți md5 sau sha256")
			return
		}

		if match {
			fmt.Println("✅ Hash-ul introdus corespunde!")
		} else {
			fmt.Println("❌ Hash-ul introdus NU corespunde!")
		}
	}
}